package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

type S3Service struct {
	Client     *s3.Client
	BucketName string
	Region     string
}

func NewS3Service(s3Client *s3.Client, cfg *config.AWSConfig) *S3Service {
	return &S3Service{
		Client:     s3Client,
		BucketName: cfg.BucketName,
		Region:     cfg.Region,
	}
}

func (s *S3Service) UploadImage(filePath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("invalid file type: only images are allowed for file %s", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		return "", fmt.Errorf("invalid file type for file %s: only image files are allowed", filePath)
	}

	uniqueFileName := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(uniqueFileName),
		Body:   bytes.NewReader(buffer),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file %s: %v", filePath, err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, s.Region, uniqueFileName)

	return fileURL, nil
}
