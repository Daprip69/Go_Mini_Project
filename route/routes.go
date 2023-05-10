package route

import (
	"net/http"
	"Go_Mini_Project/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUserController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	// book collection
	consultation := e.Group("/consultations")
	consultation.GET("", controller.GetConsultationsController)
	consultation.GET("/:id", controller.GetConsultationController)
	consultation.POST("", controller.CreateConsultationController)
	consultation.PUT("/:id", controller.UpdateConsultationController)
	consultation.DELETE("/:id", controller.DeleteConsultationController)

	// book collection
	feedback := e.Group("/feedbacks")
	feedback.GET("", controller.GetFeedbacksController)
	feedback.GET("/:id", controller.GetFeedbackController)
	feedback.POST("", controller.CreateFeedbackController)
	feedback.PUT("/:id", controller.UpdateFeedbackController)
	feedback.DELETE("/:id", controller.DeleteFeedbackController)

}
