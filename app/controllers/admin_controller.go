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
			"title": "ドキュメント一覧",
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

func (controller *AdminController) DocumentsCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.Form("title")
		slug := c.Form("slug")
		content := c.Form("content")
		controller.DB.Create(&models.Document{Title: title, Slug: slug, Text: content})
		return c.Redirect(http.StatusFound, "/admin/documents")
	}
}

func (controller *AdminController) DocumentsUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		var doc models.Document
		controller.DB.Find(&doc, c.Param("id"))
		doc.Title = c.Form("title")
		doc.Slug = c.Form("slug")
		doc.Text = c.Form("content")
		controller.DB.Save(&doc)
		return c.Redirect(http.StatusFound, "/admin/documents")
	}
}

func (controller *AdminController) DocumentsEdit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var doc models.Document
		controller.DB.Find(&doc, c.Param("id"))
		data := map[string]interface{}{
			"title": doc.Title,
			"doc":   doc,
		}
		return c.Render(http.StatusOK, "admin/documents/edit.html", data)
	}
}

func (controller *AdminController) DocumentsShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		var doc models.Document
		controller.DB.Find(&doc, c.Param("id"))
		data := map[string]interface{}{
			"title": doc.Title,
			"doc":   doc,
		}
		return c.Render(http.StatusOK, "admin/documents/show.html", data)
	}
}
