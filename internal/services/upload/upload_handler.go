package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"setUp/internal/utils"
	// "setUp/pkg/jwt"

	"setUp/internal/minio"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// "path/filepath"
)

type UploadHandler struct {
	log *zap.Logger
}

func NewUploadHandler(log *zap.Logger) *UploadHandler {
	return &UploadHandler{log: log}
}

func (h *UploadHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}
	// Simpan file ke disk (sementara)
	tempFilePath := "./" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan file sementara"})
		return
	}
	timeNow := time.Now().Unix()
	filename := fmt.Sprintf("%d-%s", timeNow, file.Filename)
	contentType := file.Header.Get("Content-Type")

	// // Implement file upload logic here
	err = minio.UploadFileInMinio(h.log, filename, tempFilePath, contentType)
	if err != nil {
		h.log.Error("Error uploading file", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}
	h.log.Info("File upload handler called")
	utils.SuccessResp(c, "success upload file", map[string]string{
		"filename": filename,
	})
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
	// ubah filename
	unixTime := time.Now().Unix()
	req.FileName = fmt.Sprintf("%d_%s", unixTime, req.FileName)

	// // Implement presigned URL generation logic here
	presignedURL, err := minio.GetPresignedURLFromMinio(h.log, req.FileName)
	if err != nil {
		h.log.Error("Error generating presigned URL", zap.Error(err))
		utils.ErrorResp(c, http.StatusInternalServerError, "Failed to generate presigned URL")
		return
	}
	MinioEndpt := os.Getenv("MINIO_ENDPOINT")
	MinioBucket := os.Getenv("MINIO_BUCKET")

	url := fmt.Sprintf(`%s/%s/%s`, MinioEndpt, MinioBucket, req.FileName)
	utils.SuccessResp(c, "Presigned URL generated successfully",
		map[string]interface{}{"presign-url": presignedURL, "filename": req.FileName, "url": url})
}

func (h *UploadHandler) GetDownloadFile(c *gin.Context) {
	// var req PresignedUrlRequest
	filename := c.Param("filename")
	if filename == "" {
		fmt.Println("Error: file_name parameter is missing")
		utils.ErrorResp(c, http.StatusBadRequest, "Invalid request")
		return
	}

	obj, err := minio.DownloadFileFromMinio(h.log, filename, "./")
	if err != nil {
		h.log.Error("Error generating File", zap.Error(err))
		utils.ErrorResp(c, http.StatusInternalServerError, "Failed to generate file")
		return
	}
	defer obj.Close()
	stat, err := obj.Stat()
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, "Failed to generate file")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", stat.ContentType)

	c.DataFromReader(http.StatusOK, stat.Size, stat.ContentType, obj, nil)

	// Stream file ke client
	//
	//	c.Stream(func(w io.Writer) bool {
	//		_, err := io.Copy(w, obj)
	//		return err == nil
	//	})
}
