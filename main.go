package main

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-TA/domain/userAggregate"
	"github.com/On-A-Rocket/Authorization-TA/domain/workAggregate"

	_ "github.com/On-A-Rocket/Authorization-TA/docs"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Authoriztion-TA Service API
// @version 1.0
// @description This is Absenteeism and tardiness management Server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9001
// @BasePath
func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	userAggregate.Initialize(e)
	workAggregate.Initialize(e)

	//work := e.Group("/work")
	//work.GET("/newworkhistory", newWorkHistory)

	// // 	기본 인증 미들웨어는 HTTP 기본 인증을 제공합니다.
	// // 유효한 자격 증명의 경우 다음 처리기를 호출합니다.
	// // 자격 증명이 없거나 유효하지 않은 경우 "401 - Unauthorized" 응답을 보냅니다.
	// g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "joe" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":9001")) // localhost:9001
}

// @Summary Insert workhistory
// @Description Insert workhistory
// @Accept json
// @Produce json
// @Param WorkhistoryBody body models.WorkHistory true "WorkHistory Info Body"
// @Success 200 {object} models.WorkHistory
// @Router /wh [post]
// func newWorkHistory(c echo.Context) (err error) {
// 	u := new(models.WorkHistory)
// 	if err = c.Bind(u); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	if err = c.Validate(u); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, u)
// }
