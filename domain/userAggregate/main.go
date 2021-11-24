package userAggregate

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo) {

	newuser := e.Group("/user")
	newuser.GET("/newuser", newUser)
}

func newUser(c echo.Context) (err error) {
	u := new(user)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}

	//var iuser iuser
	fmt.Println(NewUser(u))

	return c.JSON(http.StatusOK, "test OK")

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
