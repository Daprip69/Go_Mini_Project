package controller

import (
	"Go_Mini_Project/model"
	"Go_Mini_Project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetConsultationsController(c echo.Context) error {
	consultations, e := usecase.GetListConsultations()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       "success",
		"consultation": consultations,
	})
}

func GetConsultationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	consultation, err := usecase.GetConsultation(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get consultation",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":        "success",
		"consultations": consultation,
	})
}

// create consultation blog
func CreateConsultationController(c echo.Context) error {
	consultation := model.Consultation{}
	c.Bind(&consultation)

	if err := usecase.CreateConsultation(&consultation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create blog",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success create new consultation",
		"consultation": consultation,
	})
}

// delete consultation by id
func DeleteConsultationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteConsultation(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete consultation",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf consultation tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete consultation",
	})
}

// update consultation by id
func UpdateConsultationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	consultation := model.Consultation{}
	c.Bind(&consultation)
	consultation.ID = uint(id)
	if err := usecase.UpdateConsultation(&consultation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update consultation",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf consultation tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
	})
}
