package utils

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type QueryParams struct {
	Filters    map[string]string // e.g. map["name"] = "john"
	SortBy     string            // e.g. "created_at"
	SortOrder  string            // "asc" or "desc"
	Page       int               // page number
	PageSize   int               // items per page
}

// ApplyQuery applies filter, sort, and pagination to the given DB query.
func ApplyQuery(db *gorm.DB, params QueryParams) *gorm.DB {
	// Filters
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
	db = db.Offset(offset).Limit(params.PageSize)

	return db
}
