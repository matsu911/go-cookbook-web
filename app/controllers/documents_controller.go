package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/matsu911/go-cookbook-web/app/models"
)

type DocumentsController struct {
	DB *gorm.DB
}

func (controller *DocumentsController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		slug := c.P(0)
		var doc models.Document
		if controller.DB.Where(&models.Document{Slug: slug}).First(&doc).RecordNotFound() {
			return c.Render(http.StatusNotFound, "errors/404.html", nil)
		}
		data := map[string]interface{}{
			"title": doc.Title,
			"doc":   doc,
		}
		return c.Render(http.StatusOK, "documents/show.html", data)
	}
}
