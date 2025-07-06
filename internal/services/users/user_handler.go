package services
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"setUp/internal/utils"
	"setUp/pkg/jwt"
)

type UserHandler struct {
	uc  UserUsecase
	log *zap.Logger
}

func NewUserHandler(uc UserUsecase, log *zap.Logger) *UserHandler {
	return &UserHandler{uc: uc, log: log}
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=100"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}
	// insert log login
	h.uc.InsertLogLogin(c, req.Email, true)

	clientIP := c.ClientIP()
	// Cek apakah user sudah ada
	user, err := h.uc.GetByEmail(c, req.Email)
	if err != nil {
		h.log.Error("Error getting user by email", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	// Jika user tidak ditemukan
	if user == nil {
		h.log.Warn("Login failed: user not found", zap.String("Email", req.Email), zap.String("client_ip", clientIP))
		h.uc.InsertLogLogin(c, req.Email, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	// Cek password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		h.log.Warn("Login failed: invalid password", zap.String("Email", req.Email), zap.String("client_ip", clientIP))
		h.uc.InsertLogLogin(c, req.Email, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := jwt.CreateToken(user.ID.String(), user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	refreshToken, err := jwt.CreateRefreshToken(user.ID.String(), user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	h.log.Info("Login Success", zap.String("Email", req.Email), zap.String("client_ip", clientIP))

	user.Password = "" // Hapus password dari response
	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user": map[string]interface{}{
			"id":      user.ID,
			"email":   user.Email,
			"name":    user.Username,
			"address": user.Address,
		},
		"message":       "Login successful",
		"client_ip":     clientIP,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.uc.GetAllUsers()
	if err != nil {
		h.log.Error("Error getting all users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ErrorMessage))
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
	return
}