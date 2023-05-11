package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Counselor', 'Client');default:'Client'; not-null"`
	Token    string `gorm:"-"`
}