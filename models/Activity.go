package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type ActivityNull struct {
}
