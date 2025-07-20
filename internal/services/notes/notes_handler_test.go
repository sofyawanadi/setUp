package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
	// "setUp/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"setUp/internal/logger"
	"setUp/pkg/database"
	"log"
	"github.com/joho/godotenv"

)


func TestGetNoteByID_WithRealJWTHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logZap := logger.NewLogger()
	defer logZap.Sync()
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	noteRepo := NewNoteRepository(db, logZap)
	noteUC := NewNoteUsecase(noteRepo, logZap)
	noteHandler := NewNoteHandler(noteUC, logZap)

	router := gin.Default()
	// router.Use(middleware.AuthMiddleware()) // Auth check pakai middleware
	api := router.Group("/api/v1")
	{
		api.GET("/notes/:id", noteHandler.GetNote)
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/notes/69091edf-8e77-47fd-88de-08c2d42c8b32", nil)

	// Masukkan Authorization dan User-Agent header seperti di curl
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5heWZvc0BleGFtcGxlLmNvbSIsImV4cCI6MTc1Mjc3MTg5NywiaWF0IjoxNzUyNzY4Mjk3LCJpZCI6ImJkOGVmNmY0LWRiMWYtNDNjYy1iYmQyLTZiZGZlNDI5MzQ2MSJ9.XnJCUrFWlQGTQO5dUX95ihp2tKtAAZgnZ4M_329BPC4")
	// req.Header.Set("User-Agent", "insomnia/11.3.0")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Ubah sesuai ekspektasi jika kamu punya DB/mock
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "text28") // atau cek ID/response yang sesuai
}
