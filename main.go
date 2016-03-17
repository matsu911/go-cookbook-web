package main

import (
	"html/template"
	"io"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/matsu911/go-cookbook-web/app"
	"github.com/matsu911/go-cookbook-web/app/controllers"
	zglob "github.com/mattn/go-zglob"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := app.ConnectDB()
	// Echo instance
	e := echo.New()
	if os.Getenv("GIN_ENV") == "development" {
		e.SetDebug(true)
	}
	matches, err := zglob.Glob("app/views/**/*.html")
	e.SetRenderer(&Template{
		templates: template.Must(template.ParseFiles(matches...)),
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	adminController := &controllers.AdminController{DB: db}
	topController := &controllers.TopController{DB: db}
	documentsController := &controllers.DocumentsController{DB: db}

	// Routes
	e.Static("/assets", "public/assets")
	e.Get("/", topController.Index())
	e.Get("/documents/*", documentsController.Show())

	admin := e.Group("/admin", middleware.BasicAuth(func(usr, pwd string) bool {
		if usr == os.Getenv("BASIC_AUTH_ID") && pwd == os.Getenv("BASIC_AUTH_PASSWORD") {
			return true
		}
		return false
	}))
	admin.Get("", adminController.Index())
	admin.Get("/", adminController.Index())
	admin.Get("/documents", adminController.DocumentsIndex())
	admin.Get("/documents/new", adminController.DocumentsNew())
	admin.Post("/documents/create", adminController.DocumentsCreate())
	admin.Get("/documents/:id", adminController.DocumentsShow())
	admin.Get("/documents/:id/edit", adminController.DocumentsEdit())
	admin.Post("/documents/:id/update", adminController.DocumentsUpdate())

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	// Start server
	e.Run(standard.New(":" + port))
}
