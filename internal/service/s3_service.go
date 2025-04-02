package service

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
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

type IS3Service interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type S3Service struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Service(s3Client *s3.Client, cfg *config.AWSConfig) IS3Service {
	return &S3Service{
		Client:     s3Client,
		BucketName: cfg.BucketName,
	}
}

func (s *S3Service) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("invalid file type: only images are allowed")
	}

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	mimeType := http.DetectContentType(buffer.Bytes())
	if !strings.HasPrefix(mimeType, "image/") {
		return "", fmt.Errorf("invalid file type: only image files are allowed")
	}

	uniqueFileName := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(uniqueFileName),
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, "ap-southeast-1", uniqueFileName)

	return fileURL, nil
}
