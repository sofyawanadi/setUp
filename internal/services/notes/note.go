package services

import "setUp/internal/domain"

type Note struct {
	domain.BaseModel
	Title   string `gorm:"column:title;not null"`
	Content string `gorm:"column:content;not null"`
}

func (Note) TableName() string {
	return "notes"
}
