package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	// "setUp/internal/middleware"
	"log"
	"setUp/internal/logger"
	"setUp/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initHandler()*NoteHandler{
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

	return noteHandler
}
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5heWZvc0BleGFtcGxlLmNvbSIsImV4cCI6MTc1MzE4NjIyNSwiaWF0IjoxNzUzMTgyNjI1LCJpZCI6ImJkOGVmNmY0LWRiMWYtNDNjYy1iYmQyLTZiZGZlNDI5MzQ2MSJ9.5qYK_fXRFo-k8ez4zVeVNtEo0GKNzbi_5Gxf6MvOX_I"
const noteId = "69091edf-8e77-47fd-88de-08c2d42c8b32"
func TestGetNoteByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	noteHandler := initHandler()
	router := gin.Default()
	// router.Use(middleware.AuthMiddleware()) // Auth check pakai middleware
	api := router.Group("/api/v1")
	{
		api.GET("/notes/:id", noteHandler.GetNote)
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/notes/"+noteId, nil)

	// Masukkan Authorization dan User-Agent header seperti di curl
	req.Header.Set("Authorization", "Bearer "+token)
	// req.Header.Set("User-Agent", "insomnia/11.3.0")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Ubah sesuai ekspektasi jika kamu punya DB/mock
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "text2") // atau cek ID/response yang sesuai
}

func TestGetNoteByIDNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	noteHandler := initHandler()
	router := gin.Default()
	// router.Use(middleware.AuthMiddleware()) // Auth check pakai middleware
	api := router.Group("/api/v1")
	{
		api.GET("/notes/:id", noteHandler.GetNote)
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/notes/4632d71a-1784-4c8c-98ef-a19670fbc93d", nil)

	// Masukkan Authorization dan User-Agent header seperti di curl
	req.Header.Set("Authorization", "Bearer "+token)
	// req.Header.Set("User-Agent", "insomnia/11.3.0")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Ubah sesuai ekspektasi jika kamu punya DB/mock
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "note not found") // atau cek ID/response yang sesuai
}

func TestCreateNoteSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()

	router.POST("/api/v1/notes", noteHandler.CreateNote)

	// Body JSON
	noteJSON := `{
		"title": "Contoh Judul",
		"content": "Isi catatan di sini"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewBuffer([]byte(noteJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Cek status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Cek response body (misalnya return ID atau message)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response["data"])            // Asumsikan response mengandung ID
	assert.Equal(t, "req created successfully", response["message"]) // Atau sesuaikan dengan pesan sebenarnya
}

func TestCreateNoteValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()

	router.POST("/api/v1/notes", noteHandler.CreateNote)

	// Body JSON
	noteJSON := `{
		"title": "Contoh Judul"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewBuffer([]byte(noteJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	
	// Cek response body (misalnya return ID atau message)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	fmt.Println("err",response)
	// require.NoError(t, err)
	
	// Cek status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, response["error"])            // Asumsikan response mengandung ID
	assert.Equal(t, "validasi gagal", response["error"]) // Atau sesuaikan dengan pesan sebenarnya
}

func TestGetAllNotes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()
	router.GET("/api/v1/notes", noteHandler.GetNotes)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/notes", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	// Misalnya cek minimal 0 hasil
	data, ok := response["data"].([]interface{})
	require.True(t, ok)
	assert.GreaterOrEqual(t, len(data), 0)
	assert.NotEmpty(t, response["pagination"])    
}

func TestUpdateNoteByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()
	router.PUT("/api/v1/notes", noteHandler.UpdateNote)

	updateJSON := `{
		"id":"`+noteId+`",
		"title": "Judul Diperbarui",
		"content": "Isi diperbarui"
	}`
	fmt.Println("json",updateJSON)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/notes", bytes.NewBuffer([]byte(updateJSON)))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response["data"])
	assert.Equal(t, "req updated successfully", response["message"]) // Sesuaikan isi responsenya
}

func TestUpdateNoteByIDValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()
	router.PUT("/api/v1/notes", noteHandler.UpdateNote)

	updateJSON := `{
		"id":"`+noteId+`",
		"title": "Judul Diperbarui"
	}`
	fmt.Println("json",updateJSON)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/notes", bytes.NewBuffer([]byte(updateJSON)))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	 json.Unmarshal(w.Body.Bytes(), &response)
	// Cek status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, response["error"])            // Asumsikan response mengandung ID
	assert.Equal(t, "validasi gagal", response["error"])  // Sesuaikan isi responsenya
}

func TestDeleteNoteByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteHandler := initHandler()
	router := gin.Default()
	router.DELETE("/api/v1/notes/:id", noteHandler.DeleteNote)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/notes/"+noteId, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "note deleted successfully", response["message"]) // Sesuaikan isi response
}
