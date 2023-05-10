package database

import (
	"Go_Mini_Project/config"
	"Go_Mini_Project/model"
)

func CreateConsultation(consultation *model.Consultation) error {
	if err := config.DB.Create(consultation).Error; err != nil {
		return err
	}
	return nil
}

func GetConsultations() (consultations []model.Consultation, err error) {
	if err = config.DB.Find(&consultations).Error; err != nil {
		return
	}
	return
}

func GetConsultation(id uint) (consultation model.Consultation, err error) {
	consultation.ID = id
	if err = config.DB.First(&consultation).Error; err != nil {
		return
	}
	return
}

func GetConsultationsByUserId(userID uint) (consultation model.Consultation, err error) {
	consultation.UserID = userID
	if err = config.DB.Find(&consultation).Error; err != nil {
		return
	}
	return
}

func UpdateConsultation(consultation *model.Consultation) error {
	if err := config.DB.Updates(consultation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteConsultation(consultation *model.Consultation) error {
	if err := config.DB.Delete(consultation).Error; err != nil {
		return err
	}
	return nil
}
