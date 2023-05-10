package usecase

import (
	"Go_Mini_Project/model"
	"Go_Mini_Project/repository/database"
	"errors"
	"fmt"
)

func CreateFeedback(feedback *model.Feedback) error {

	// check title cannot be empty
	if feedback.Description == "" {
		return errors.New("feedback description cannot be empty")
	}

	// check creator
	if feedback.Status == "" {
		return errors.New("feedback status cannot be empty")
	}

	err := database.CreateFeedback(feedback)
	if err != nil {
		return err
	}

	return nil
}

func GetFeedback(id uint) (feedback model.Feedback, err error) {
	feedback, err = database.GetFeedback(id)
	if err != nil {
		fmt.Println("GetFeedback: Error getting feedback from database")
		return
	}
	return
}

func GetListFeedbacks() (feedbacks []model.Feedback, err error) {
	feedbacks, err = database.GetFeedbacks()
	if err != nil {
		fmt.Println("GetListFeedback: Error getting feedbacks from database")
		return
	}
	return
}

func UpdateFeedback(feedback *model.Feedback) (err error) {
	err = database.UpdateFeedback(feedback)
	if err != nil {
		fmt.Println("UpdateFeedback : Error updating feedback, err: ", err)
		return
	}

	return
}

func DeleteFeedback(id uint) (err error) {
	feedback := model.Feedback{}
	feedback.ID = id
	err = database.DeleteFeedback(&feedback)
	if err != nil {
		fmt.Println("DeleteFeedback : error deleting feedback, err: ", err)
		return
	}

	return
}
