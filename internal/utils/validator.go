package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var validate = validator.New()
type ValidationErrorResponse struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details"`
}
func ValidateRequest(req interface{}, c *gin.Context, log *zap.Logger) bool {
	err := validate.Struct(req)
	if err == nil {
		return true
	}

	// Kumpulkan error dalam map
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		errors[field] = fmt.Sprintf("Field '%s' tidak valid (rule: %s)", field, tag)
	}

	// Marshal error ke JSON string untuk logging
	jsonErr, _ := json.MarshalIndent(errors, "", "  ")
	log.Warn("Validasi gagal", zap.String("errors", string(jsonErr)))

	// Kirim response ke client
	resp := ValidationErrorResponse{
		Error:   "validasi gagal",
		Details: errors, // bisa berupa map[string]string atau []string tergantung isinya
	}
	c.JSON(http.StatusBadRequest, resp)

	return false
}
