package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo"
)

// Get all book controller
func GetBookController(c echo.Context) error {
	books, err := database.GetBook()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all book",
		Data:    books,
	})
}

// Create book controller
func CreateBookController(c echo.Context) error {
	books := models.Book{}
	c.Bind(&books)

	books, err := database.CreateBook(books)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create book",
		Data:    books,
	})
}

// Get book by id controller
func GetBookByIdController(c echo.Context) error {
	BookId := c.Param("id")

	books, err := database.GetBookById(BookId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get book by id",
		Data:    books,
	})
}

// Update book by id controller
func UpdateBookByIdController(c echo.Context) error {
	BookId := c.Param("id")

	books := models.Book{}
	c.Bind(&books)

	books, err := database.UpdateBookById(books, BookId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update book by id",
		Data:    books,
	})
}

// Delete book by id controller
func DeleteBookByIdController(c echo.Context) error {
	BookId := c.Param("id")

	_, err := database.DeleteBook(BookId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete book by id",
	})

}