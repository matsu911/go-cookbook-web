package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type AdminController struct{}

func (AdminController) DocumentsNew() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "documents/new.html", nil)
	}
}

func (AdminController) DocumentsShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "documents/show.html", nil)
	}
}
