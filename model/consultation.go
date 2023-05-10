package model

import "gorm.io/gorm"

type Consultation struct {
	gorm.Model
	Name     		string `json:"name" form:"name"`
	Symptoms 		string `json:"symptoms" form:"symptoms"`
	Date    		string `json:"date" form:"date"`
	UserID   		uint   `json:"user_id" form:"user_id"`
	ConsultationID  uint   `json:"consultation_id" form:"consultation_id"`

}
