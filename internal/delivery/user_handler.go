// user_handler.go
package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"setUp/internal/usecase"
)

type UserHandler struct {
	uc  *usecase.UserUsecase
	log *zap.Logger
}

func NewUserHandler(uc *usecase.UserUsecase, log *zap.Logger) *UserHandler {
	return &UserHandler{uc: uc, log: log}
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=100"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		// Kumpulkan error ke dalam map
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()

			// Bisa disesuaikan untuk pesan error yang lebih user-friendly
			errors[field] = fmt.Sprintf("Field '%s' tidak valid (rule: %s)", field, tag)
		}

		// Misal ingin kirim sebagai JSON response
		jsonErr, _ := json.MarshalIndent(errors, "", "  ")

		h.log.Warn("Validasi gagal", zap.String("errors", string(jsonErr)))
		h.log.Error("Error marshalling JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "validasi gagal"})
		return

	}
    clientIP := c.ClientIP()
	token, err := h.uc.Login(req.Username, req.Password)
	if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
    h.log.Info("Login Success",zap.String("username", req.Username), zap.String("client_ip", clientIP))

	c.JSON(http.StatusOK, gin.H{"token": token})
}
