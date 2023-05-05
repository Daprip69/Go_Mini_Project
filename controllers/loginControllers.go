package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/middleware"
	"praktikum/models"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	user.Token, err = middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create user",
		Data:    user,
	})
}