package utils

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SuccessMessage = "Operation completed successfully"
const ErrorMessage = "An error occurred during the operation"

type PaginationMeta struct {
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SuccessResponse adalah response dengan data
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessWithPaginationResponse adalah response dengan data + pagination
type SuccessWithPaginationResponse struct {
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

func SuccessResp(c *gin.Context, message string, data interface{}) {
	resp := SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}
func SuccessWithPaginationResp(c *gin.Context, message string, data interface{}, page, pageSize, total int64) {
	totalPages := int64(math.Ceil(float64(total) / float64(pageSize)))
	resp := SuccessWithPaginationResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: PaginationMeta{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func ErrorResp(c *gin.Context, status int, message string) {
	resp := BaseResponse{
		Success: false,
		Message: message,
	}
	c.JSON(status, resp)
}
