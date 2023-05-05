package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo"
)

func GetBlogController(c echo.Context) error {
	blogs, err := database.Getblog()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all blog",
		Data:    blogs,
	})
}

func CreateBlogController(c echo.Context) error {
	blogs := models.Blog{}
	c.Bind(&blogs)

	blogs, err := database.CreateBlog(blogs)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create blog",
		Data:    blogs,
	})

}

func GetBlogByIdController(c echo.Context) error {
	BlogId := c.Param("id")

	blogs, err := database.GetBlogById(BlogId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get blog by id",
		Data:    blogs,
	})
}

func UpdateBlogByIdController(c echo.Context) error {
	BlogId := c.Param("id")
	blogs := models.Blog{}
	c.Bind(&blogs)

	blogs, err := database.UpdateBlog(blogs, BlogId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update blog by id",
		Data:    blogs,
	})
}

func DeleteBlogByIdController(c echo.Context) error {
	BlogId := c.Param("id")

	blogs, err := database.DeleteBlog(BlogId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete blog by id",
		Data:    blogs,
	})
}