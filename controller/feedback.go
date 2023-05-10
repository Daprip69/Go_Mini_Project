package controller

import (
	"net/http"
	"Go_Mini_Project/model"
	"Go_Mini_Project/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetFeedbacksController(c echo.Context) error {
	feedbacks, e := usecase.GetListFeedbacks()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"feedback":  feedbacks,
	})
}

func GetFeedbackController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	feedback, err := usecase.GetFeedback(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get feedback",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"feedback":  feedback,
	})
}

// create new feedback
func CreateFeedbackController(c echo.Context) error {
	feedback := model.Feedback{}
	c.Bind(&feedback)

	if err := usecase.CreateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"feedback":    feedback,
	})
}

// delete feedback by id
func DeleteFeedbackController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteFeedback(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete feedback",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf feedback tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update feedback by id
func UpdateFeedbackController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	feedback := model.Feedback{}
	c.Bind(&feedback)
	feedback.ID = uint(id)
	if err := usecase.UpdateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update feedback",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf feedback tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update feedback",
	})
}
