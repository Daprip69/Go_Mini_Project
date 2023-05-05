package routes

import (
	"praktikum/constants"
	"praktikum/controllers"
	"praktikum/middleware"

	mid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// Crate echo intance
	e := echo.New()

	// Midlleware
	middleware.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// User Controller
	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers, mid.JWT([]byte(constants.SECRET_JWT)))
	u.POST("", controllers.CreateUserController)
	u.GET("/:id", controllers.GetUserByIdController, mid.JWT([]byte(constants.SECRET_JWT)))
	u.PUT("/:id", controllers.UpdateUserByIdController, mid.JWT([]byte(constants.SECRET_JWT)))
	u.DELETE("/:id", controllers.DeleteUserByIdController, mid.JWT([]byte(constants.SECRET_JWT)))
	u.POST("/login", controllers.LoginUserController)

	// Book Controller
	bo := e.Group("/books")
	bo.GET("", controllers.GetBookController, mid.JWT([]byte(constants.SECRET_JWT)))
	bo.POST("", controllers.CreateBookController, mid.JWT([]byte(constants.SECRET_JWT)))
	bo.GET("/:id", controllers.GetBookByIdController, mid.JWT([]byte(constants.SECRET_JWT)))
	bo.PUT("/:id", controllers.UpdateBookByIdController, mid.JWT([]byte(constants.SECRET_JWT)))
	bo.DELETE("/:id", controllers.DeleteBookByIdController, mid.JWT([]byte(constants.SECRET_JWT)))

	// Blog Controller
	bl := e.Group("/blogs")
	bl.GET("", controllers.GetBlogController)
	bl.POST("", controllers.CreateBlogController)
	bl.GET("/:id", controllers.GetBlogByIdController)
	bl.PUT("/:id", controllers.UpdateBlogByIdController)
	bl.DELETE("/:id", controllers.DeleteBlogByIdController)

	return e
}