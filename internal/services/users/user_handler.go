package services

import (
	"fmt"
	"net/http"

	"setUp/internal/utils"
	"setUp/pkg/jwt"

	"github.com/gin-gonic/gin"
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

// GetUser godoc
// @Summary Get a user by ID
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
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
	user, err := h.uc.InsertUser(c,req)
	if err != nil {
		h.log.Error("Error getting user by ID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(),"success":false})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	user.Password = "" // Hapus password dari response
	utils.SuccessResp(c, "User retrieved successfully", map[string]interface{}{
		"data":user,
		"message":"success create user",
		"success":true,
	})
	return 
}