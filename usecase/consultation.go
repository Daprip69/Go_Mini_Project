package usecase

import (
	"errors"
	"fmt"
	"Go_Mini_Project/model"
	"Go_Mini_Project/repository/database"
)

type ConsultationUsecase interface {
	CreateConsultation(consultation *model.Consultation) error
	GetConsultation(id uint) (consultation model.Consultation, err error)
	GetListConsultations() (consultations []model.Consultation, err error)
	UpdateConsultation(consultation *model.Consultation) (err error)
	DeleteConsultation(id uint) (err error)
}

func CreateConsultation(consultation *model.Consultation) error {

	// check name cannot be empty
	if consultation.Name == "" {
		return errors.New("consultation Name cannot be empty")
	}

	// check symptoms cannot be empty
	if consultation.Symptoms == "" {
		return errors.New("consultation Symptoms cannot be empty")
	}

	err := database.CreateConsultation(consultation)
	if err != nil {
		return err
	}

	return nil
}

func GetConsultation(id uint) (consultation model.Consultation, err error) {
	consultation, err = database.GetConsultation(id)
	if err != nil {
		fmt.Println("GetBlog: Error getting blog from database")
		return
	}
	return
}

func GetListConsultations() (consultations []model.Consultation, err error) {
	consultations, err = database.GetConsultations()
	if err != nil {
		fmt.Println("GetListConsultations: Error getting consultations from database")
		return
	}
	return
}

func UpdateConsultation(consultation *model.Consultation) (err error) {
	err = database.UpdateConsultation(consultation)
	if err != nil {
		fmt.Println("UpdateConsultation : Error updating consultation, err: ", err)
		return
	}

	return
}

func DeleteConsultation(id uint) (err error) {
	consultation := model.Consultation{}
	consultation.ID = id
	err = database.DeleteConsultation(&consultation)
	if err != nil {
		fmt.Println("DeleteConsultation : error deleting consultation, err: ", err)
		return
	}

	return
}
