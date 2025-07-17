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

func SuccessResp(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SuccessWithPaginationResp(c *gin.Context, message string, data interface{}, page, pageSize, total int64) {
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
		"pagination": PaginationMeta{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: int64(totalPages),
		},
	})
}

func ErrorResp(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
}
