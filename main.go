package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/On-A-Rocket/Authorization-TA/docs"
	"github.com/On-A-Rocket/Authorization-TA/model"
	"gopkg.in/yaml.v3"

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

// @host localhost:1323
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323
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
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func getUser(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// db, err := sql.Open("mysql", "root:dkssud123@tcp(127.0.0.1:3306)/test")
// if err != nil {
// 			fmt.Println(err.Error())
// } else {
// 			fmt.Println("db is connected")
// }
// defer db.Close()
// // make sure connection is available
// err = db.Ping()
// if err != nil {
// 			fmt.Println(err.Error())
// }

// func getUser(c echo.Context) error {

// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)

// }
