package services

import (
	"fmt"
	"net/http"

	"setUp/internal/utils"
	// "setUp/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"setUp/internal/minio"
	// "path/filepath"
)

type UploadHandler struct {
	log *zap.Logger
}

func NewUploadHandler(log *zap.Logger) *UploadHandler {
	return &UploadHandler{log: log}
}

type UploadRequest struct {
	FileName string `json:"file_name" form:"file_name" validate:"required"`
	File     []byte `json:"file" form:"file" validate:"required"`
}

func (h *UploadHandler) UploadFile(c *gin.Context) {
	var req UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}
	// contentType := http.DetectContentType(req.File)
	// // // Implement file upload logic here
	// err := minio.UploadFileInMinio(h.log, req.FileName, req.File , contentType)
	// if err != nil {
	// 	h.log.Error("Error uploading file", zap.Error(err))
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
	// 	return
	// }
	h.log.Info("File upload handler called")
	c.JSON(200, gin.H{"message": "File uploaded successfully"})
}

type PresignedUrlRequest struct {
	FileName string `json:"file_name" form:"file_name" validate:"required"`
}

func (h *UploadHandler) GetPresignedUrl(c *gin.Context) {
	var req PresignedUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		utils.ErrorResp(c, http.StatusBadRequest, "Invalid request")
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}

	// // Implement presigned URL generation logic here
	presignedURL, err := minio.GetPresignedURLFromMinio(h.log, req.FileName)
	if err != nil {
		h.log.Error("Error generating presigned URL", zap.Error(err))
		utils.ErrorResp(c, http.StatusInternalServerError, "Failed to generate presigned URL")
		return
	}
	utils.SuccessResp(c, "Presigned URL generated successfully",
		map[string]interface{}{"url": presignedURL})
}
