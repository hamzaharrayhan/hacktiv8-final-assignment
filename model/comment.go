package model

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Message   string    `json:"message" gorm:"not null"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
