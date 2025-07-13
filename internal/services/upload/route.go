package services
import (
	"github.com/gin-gonic/gin"
	"setUp/internal/middleware"
)

func RouteUpload(rg *gin.RouterGroup, uploadHandler *UploadHandler) {
	userGroup := rg.Group("/upload").Use(middleware.AuthMiddleware()) 
	{
		userGroup.POST("/", uploadHandler.UploadFile)
		userGroup.GET("/:filename", uploadHandler.GetDownloadFile)
		userGroup.GET("/get-url", uploadHandler.GetPresignedUrl)
	}
}