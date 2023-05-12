package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint           `json:"id" form:"id" gorm:"primary_key"`
	Name         string         `json:"name" form:"name"`
	Email        string         `json:"email" form:"email"`
	Password     string         `json:"password" form:"password"`
	Role         string         `json:"role" form:"role" gorm:"type:enum('Counselor', 'Client');default:'Client'; not-null"`
	Token        string         `gorm:"-"`
	Consultations []Consultation `gorm:"foreignKey:UserID"`
}