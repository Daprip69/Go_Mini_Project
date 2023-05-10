package database

import (
	"Go_Mini_Project/config"
	"Go_Mini_Project/model"
)

func CreateFeedback(feedback *model.Feedback) error {
	if err := config.DB.Create(feedback).Error; err != nil {
		return err
	}
	return nil
}

func GetFeedbacks() (feedbacks []model.Feedback, err error) {
	if err = config.DB.Find(&feedbacks).Error; err != nil {
		return
	}
	return
}

func GetFeedback(id uint) (feedback model.Feedback, err error) {
	feedback.ID = id
	if err = config.DB.First(&feedback).Error; err != nil {
		return
	}
	return
}

func UpdateFeedback(feedback *model.Feedback) error {
	if err := config.DB.Updates(feedback).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFeedback(feedback *model.Feedback) error {
	if err := config.DB.Delete(feedback).Error; err != nil {
		return err
	}
	return nil
}
