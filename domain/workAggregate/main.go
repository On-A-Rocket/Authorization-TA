package workAggregate

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo) {

	work := e.Group("/work")
	work.GET("/newworkhistory", newWorkHistory)
}

// @Summary Insert workhistory
// @Description Insert workhistory
// @Accept json
// @Produce json
// @Param WorkhistoryBody body models.WorkHistory true "WorkHistory Info Body"
// @Success 200 {object} models.WorkHistory
// @Router /wh [post]
func newWorkHistory(c echo.Context) (err error) {
	u := new(WorkHistory)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
