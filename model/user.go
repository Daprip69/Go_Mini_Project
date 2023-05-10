package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	UserID   int    `json:"user_id" form:"user_id" gorm:"unique"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Counselor', 'Client');default:'Client'; not-null"`
	Token    string `gorm:"-"`
}