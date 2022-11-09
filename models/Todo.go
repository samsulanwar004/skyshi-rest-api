package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              int             `json:"id"`
	ActivityGroupID int             `json:"activity_group_id"`
	Title           string          `json:"title"`
	IsActive        bool            `json:"is_active"`
	Priority        string          `json:"priority"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
}

type TodoNull struct {
}
