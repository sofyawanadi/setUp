package services

import (
	"setUp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RouteUpload(rg *gin.RouterGroup, uploadHandler *UploadHandler) {
	userGroup := rg.Group("/upload").Use(middleware.AuthMiddleware())
	{
		userGroup.POST("/", uploadHandler.UploadFile)
		userGroup.GET("/:filename", uploadHandler.GetDownloadFile)
		userGroup.GET("/get-url", uploadHandler.GetPresignedUrl)
	}
}
