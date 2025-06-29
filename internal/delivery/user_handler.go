// user_handler.go
package delivery

import (
	"fmt"
	"net/http"

	"setUp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
    return &UserHandler{uc: uc}
}

type LoginRequest struct {
    Username string `json:"username" form:"username" `
    Password string `json:"password" form:"password" `
}

func (h *UserHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.uc.Login(req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
