package model

import "time"

type SocialMedia struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"not null"`
	URL       string    `json:"social_media_url" gorm:"not null"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
