package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetNoteByID_Success(t *testing.T) {
	// Set Gin ke mode test
	gin.SetMode(gin.TestMode)

	// Buat router dan binding handler
	r := gin.Default()
	noteHandler := &NoteHandler{}
	r.GET("api/v1/notes/:id", noteHandler.GetNote)

	// Buat request ke endpoint
	req, _ := http.NewRequest(http.MethodGet, "/notes/69091edf-8e77-47fd-88de-08c2d42c8b32", nil)
	w := httptest.NewRecorder()

	// Jalankan
	r.ServeHTTP(w, req)

	// Assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Note")
}

func TestGetNoteByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	noteHandler := &NoteHandler{}
	r.GET("api/v1/notes/:id", noteHandler.GetNote)

	req, _ := http.NewRequest(http.MethodGet, "/notes/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Note not found")
}