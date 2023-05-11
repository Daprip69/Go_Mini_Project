package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Name     		string `json:"name" form:"name"`
	FeedbackID 		uint   `json:"feedback_id" form:"feedback_id"`
	Date  			string `json:"date" form:"date"`
	Description  	string `json:"description" form:"description"`
	Status  		string `json:"status" form:"status"`
	ConsultationID 	uint   `json:"consultation_id" form:"consultation_id"`
}