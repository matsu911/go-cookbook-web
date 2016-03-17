package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/matsu911/go-cookbook-web/app/models"
)

type TopController struct {
	DB *gorm.DB
}

func (controller *TopController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		var doc models.Document
		controller.DB.Where(&models.Document{Slug: "/"}).First(&doc)
		data := map[string]interface{}{
			"title": doc.Title,
			"doc":   doc,
		}
		return c.Render(http.StatusOK, "index.html", data)
	}
}
