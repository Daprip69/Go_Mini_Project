package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Date  			string `json:"date" form:"date"`
	Description  	string `json:"description" form:"description"`
	Status  		string `json:"status" form:"status"`
	ConsultationID 	uint   `json:"consultation_id" form:"consultation_id"`
}