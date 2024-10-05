package models

import (
	"time"
	// "gorm.io/gorm"
)

// Comment model
type Comment struct {
	// gorm.Model

	ID         uint      `json:"id" gorm:"primaryKey"`
	PostID     uint      `json:"post_id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
