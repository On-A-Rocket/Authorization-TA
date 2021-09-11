package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/On-A-Rocket/Authorization-TA/docs"
	"github.com/On-A-Rocket/Authorization-TA/models"
	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"

	//middleware "github.com/labstack/echo/middleware"
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

// @host localhost:7101
// @BasePath
func main() {

	var c ConnectionInfo
	c.getConnectionInfo()
	fmt.Println(c)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/user", getUser)
	e.POST("/wh", postWorkHistory)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":7101")) // localhost:1323

	// e := echo.New()
	// e.GET("/request", func(c echo.Context) error {
	// 	req := c.Request()
	// 	format := `
	// 		<code>
	// 			Protocol: %s<br>
	// 			Host: %s<br>
	// 			Remote Address: %s<br>
	// 			Method: %s<br>
	// 			Path: %s<br>
	// 		</code>
	// 	`
	// 	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	// })
	// e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))

}

type ConnectionInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Id       string `yaml:"id"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}

func (c *ConnectionInfo) getConnectionInfo() *ConnectionInfo {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body models.User true "User Info Body"
// @Success 200 {object} models.User
// @Router /user [post]
func getUser(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// @Summary Insert workhistory
// @Description Insert workhistory
// @Accept json
// @Produce json
// @Param WorkhistoryBody body models.WorkHistory true "WorkHistory Info Body"
// @Success 200 {object} models.WorkHistory
// @Router /wh [post]
func postWorkHistory(c echo.Context) (err error) {
	u := new(models.WorkHistory)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
