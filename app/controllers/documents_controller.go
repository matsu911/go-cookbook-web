package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type DocumentsController struct{}

func (DocumentsController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := map[string]string{
			"title": "test",
		}
		return c.Render(http.StatusOK, "home/index.html", data)
	}
}

func (DocumentsController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "documents/show.html", nil)
	}
}
