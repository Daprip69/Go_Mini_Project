package model

import "gorm.io/gorm"

type Consultation struct {
	gorm.Model
	ID           uint      `json:"id" form:"id" gorm:"primary_key"`
	Name         string    `json:"name" form:"name"`
	Symptoms     string    `json:"symptoms" form:"symptoms"`
	Date         string    `json:"date" form:"date"`
	UserID       uint      `json:"user_id" form:"user_id"`
	User         User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Feedbacks    []Feedback `gorm:"foreignKey:ConsultationID"`
}