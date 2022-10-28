package model

import "time"

type Photo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" gorm:"not null"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
