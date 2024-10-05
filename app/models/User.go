package models

import (
	"time"
	// "gorm.io/gorm"
)

// User model
type User struct {
	// gorm.Model

	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type UserRegister struct {
	// gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
