package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IS3Controller interface {
	UploadFile(c *gin.Context)
}

type S3Controller struct {
	s3Service service.IS3Service
}

func NewS3Controller(s service.IS3Service) IS3Controller {
	return &S3Controller{
		s3Service: s,
	}
}

func (uc *S3Controller) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest, // Corrected status code
			Error:   "File not received",
			Message: err.Error(),
		})
		return
	}
	defer file.Close()

	// Upload the file using the service
	fileURL, err := uc.s3Service.UploadFile(file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Can't Upload File",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Image Uploaded",
		Data:    fileURL,
	})
}
