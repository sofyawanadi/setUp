package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryParams struct {
	Filters   map[string]string // e.g. map["name"] = "john"
	SortBy    string            // e.g. "created_at"
	SortOrder string            // "asc" or "desc"
	Page      int64             // page number
	PageSize  int64             // items per page
}

// ApplyQuery applies filter, sort, and pagination to the given DB query.
func ApplyQuery(db *gorm.DB, params QueryParams) *gorm.DB {
	// Filters
	log.Print("Applying filters: ", params.Filters)
	for field, value := range params.Filters {
		// Basic LIKE filter, can be adjusted for exact match or numeric, etc.
		if strings.Contains(value, "%") {
			db = db.Where(fmt.Sprintf("%s LIKE ?", field), value)
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}

	// Sorting
	if params.SortBy != "" {
		order := "asc"
		if strings.ToLower(params.SortOrder) == "desc" {
			order = "desc"
		}
		db = db.Order(fmt.Sprintf("%s %s", params.SortBy, order))
	}

	// Pagination
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}
	offset := (params.Page - 1) * params.PageSize
	db = db.Offset(int(offset)).Limit(int(params.PageSize))
	return db
}

func GetFilter(c *gin.Context) map[string]string {
	filters := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		// Lewati parameter yang sudah digunakan untuk sorting dan pagination
		if key == "sort_by" || key == "sort_order" || key == "page" || key == "page_size" {
			continue
		}
		// Ambil nilai pertama dari slice values
		if len(values) > 0 && values[0] != "" {
			// Contoh: untuk pencarian nama bisa gunakan LIKE
			if key == "name" {
				filters[key] = "%" + values[0] + "%"
			} else {
				filters[key] = values[0]
			}
		}
	}
	return filters
}
