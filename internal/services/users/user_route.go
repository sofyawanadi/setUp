package services
import (
	"github.com/gin-gonic/gin"
	"setUp/internal/middleware"
)

func RouteUser(rg *gin.RouterGroup, userHandler *UserHandler) {
	rg.POST("/login", userHandler.Login)
	userGroup := rg.Group("/users").Use(middleware.AuthMiddleware()) 
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.GET("/get", userHandler.GetByID)
		userGroup.POST("/",userHandler.PostUser)
	}
}