package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/matsu911/go-cookbook-web/app/models"
)

type AdminController struct {
	DB *gorm.DB
}

func (*AdminController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/index.html", nil)
	}
}

func (controller *AdminController) DocumentsIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		var documents []models.Document
		controller.DB.Find(&documents)
		data := map[string]interface{}{
			"title": "aaaa",
			"docs":  documents,
		}
		return c.Render(http.StatusOK, "admin/documents/index.html", data)
	}
}

func (*AdminController) DocumentsNew() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/documents/new.html", nil)
	}
}

func (*AdminController) DocumentsCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/admin/documents")
	}
}

func (*AdminController) DocumentsShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/documents/show.html", nil)
	}
}
