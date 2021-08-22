package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/wlsgh199/TimeAndAttendance/docs"
)

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// // 첫 화면
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	e.POST("/user", getUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func getUser(c echo.Context) error {
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

// func getUser(c echo.Context) error {

// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)

// }
