package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	ID             uint         `json:"id" form:"id" gorm:"primary_key"`
	Date           string       `json:"date" form:"date"`
	Description    string       `json:"description" form:"description"`
	Status         string       `json:"status" form:"status"`
	ConsultationID uint         `json:"consultation_id" form:"consultation_id"`
	Consultation   Consultation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}