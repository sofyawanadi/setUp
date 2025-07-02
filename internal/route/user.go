package route

import (
	"github.com/gin-gonic/gin"
	"setUp/internal/delivery"
)

func RouteUser(rg *gin.RouterGroup, userHandler *delivery.UserHandler) {
	// User routes
	rg.POST("/login", userHandler.Login)
	userGroup := rg.Group("/user") // Apply authentication middleware to user routes
	{
		userGroup.GET("/get", userHandler.GetAllUsers)

		// Add more routes as needed
		// r.GET("/some-other-route", someOtherHandler.SomeMethod)
	}
}