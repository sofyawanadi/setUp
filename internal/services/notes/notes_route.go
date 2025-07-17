package services

import (
	"setUp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RouteNote(rg *gin.RouterGroup, noteHandler *NoteHandler) {
	noteGroup := rg.Group("/notes").Use(middleware.AuthMiddleware())
	{
		noteGroup.GET("/:id", noteHandler.GetNote)
		noteGroup.GET("/", noteHandler.GetNotes)
		noteGroup.POST("/", noteHandler.CreateNote)
		noteGroup.PUT("/", noteHandler.UpdateNote)
		noteGroup.DELETE("/:id", noteHandler.DeleteNote)
	}
}
