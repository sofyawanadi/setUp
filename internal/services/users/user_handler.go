package services

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"setUp/internal/utils"
	jwtPkg "setUp/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	// "path/filepath"
)

type UserHandler struct {
	uc  UserUsecaseInterface
	log *zap.Logger
}

func NewUserHandler(uc UserUsecaseInterface, log *zap.Logger) *UserHandler {
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

	token, err := jwtPkg.CreateToken(user.ID.String(), user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	refreshToken, err := jwtPkg.CreateRefreshToken(user.ID.String(), user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	h.log.Info("Login Success", zap.String("Email", req.Email), zap.String("client_ip", clientIP))

	utils.SendMail([]string{req.Email}, "Login Notification", "login_notification.html", map[string]interface{}{
		"Username": user.Username,
		"Email":    user.Email,
		"ClientIP": clientIP,
	})

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
		"message":   "Login successful",
		"client_ip": clientIP,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	filters := utils.GetFilter(c)
	params := utils.QueryParams{
		Filters:   filters,
		SortBy:    c.DefaultQuery("sort_by", "created_at"),
		SortOrder: c.DefaultQuery("sort_order", "desc"),
		Page:      int64(utils.ParseInt(c.DefaultQuery("page", "1"))),
		PageSize:  int64(utils.ParseInt(c.DefaultQuery("page_size", "10"))),
	}
	users, totalData, err := h.uc.GetAllUsers(c, params)
	if err != nil {
		h.log.Error("Error getting all users", zap.Error(err))
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessWithPaginationResp(c, "Users retrieved successfully", users, params.Page, params.PageSize, totalData)
	return
}

func (h *UserHandler) GetByID(c *gin.Context) {
	user, err := h.uc.GetByID(c)
	if err != nil {
		h.log.Error("Error getting user by ID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	user.Password = "" // Hapus password dari response
	utils.SuccessResp(c, "User retrieved successfully", user)
	return
}

func (h *UserHandler) PostUser(c *gin.Context) {
	var req PostUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}
	h.log.Info("[Post][users][PostUser]", zap.Any("request", req))
	user, err := h.uc.InsertUser(c, req)
	if err != nil {
		h.log.Error("Error getting user by ID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	user.Password = "" // Hapus password dari response
	utils.SuccessResp(c, "User retrieved successfully", map[string]interface{}{
		"data":    user,
		"message": "success create user",
		"success": true,
	})
	return
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	// cek apakah ada token
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// pastikan algoritma cocok
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization"})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// log.Print("Authenticated user ID:", claims)
		c.Set("userID", claims["id"])
		c.Set("exp", claims["exp"])
		c.Set("iat", claims["iat"])
		c.Set("email", claims["email"])
	}
	// Cek apakah user sudah ada
	clientIP := c.ClientIP()
	user, err := h.uc.GetByID(c)
	if err != nil {
		h.log.Error("Error getting user by email", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	// Jika user tidak ditemukan
	if user == nil {
		h.log.Warn("Login failed: user not found", zap.String("Email", user.Email), zap.String("client_ip", clientIP))
		h.uc.InsertLogLogin(c, user.Email, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	userId := c.GetString("userID")
	tokenString, err = jwtPkg.CreateToken(userId, user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	refreshToken, err := jwtPkg.CreateRefreshToken(userId, user.Email)
	if err != nil {
		h.log.Error("Error creating token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	h.log.Info("Login Success", zap.String("Email", user.Email), zap.String("client_ip", clientIP))

	utils.SuccessResp(c, "Success Refresh Token", map[string]interface{}{
		"token":         tokenString,
		"refresh_token": refreshToken,
	})
	return
}
