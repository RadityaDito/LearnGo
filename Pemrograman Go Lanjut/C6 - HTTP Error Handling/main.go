package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18,lte=100"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	// Set up the custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Custom HTTP error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	// e.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	report, ok := err.(*echo.HTTPError)
	// 	if !ok {
	// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}

	// 	c.Logger().Error(report)
	// 	c.JSON(report.Code, report)
	// }

	// e.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	report, ok := err.(*echo.HTTPError)
	// 	if !ok {
	// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// 	}

	// 	errPage := fmt.Sprintf("%d.html", report.Code)
	// 	if err := c.File(errPage); err != nil {
	// 		c.HTML(report.Code, "Errrrooooorrrrr")
	// 	}
	// }

	// Routes
	e.GET("/", handleHome)
	e.GET("/error", handleError)
	e.POST("/users", handleCreateUser)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the API!")
}

func handleError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "This is a test error")
}

func handleCreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
