package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type TopController struct{}

func (TopController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := map[string]string{
			"title": "test",
		}
		return c.Render(http.StatusOK, "home/index.html", data)
	}
}
