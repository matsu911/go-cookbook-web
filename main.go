package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/matsu911/go-cookbook-web/app"
	"github.com/matsu911/go-cookbook-web/app/controllers"
	"github.com/russross/blackfriday"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func hello2() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.HTML(http.StatusOK, string(blackfriday.MarkdownBasic(([]byte)("# test"))))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.ConnectDB()
	// Echo instance
	e := echo.New()
	e.SetRenderer(&Template{
		templates: template.Must(template.ParseGlob("views/**/*.html")),
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	adminController := new(controllers.AdminController)
	topController := new(controllers.TopController)
	documentsController := new(controllers.DocumentsController)

	// Routes
	e.Static("/assets", "public/assets")
	e.Get("/", topController.Index())
	e.Get("/documents/:id", documentsController.Show())

	admin := e.Group("/admin", middleware.BasicAuth(func(usr, pwd string) bool {
		if usr == "joe" && pwd == "secret" {
			return true
		}
		return false
	}))
	admin.Get("/documents/new", adminController.DocumentsNew())
	admin.Get("/documents/:id", adminController.DocumentsShow())

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	// Start server
	e.Run(standard.New(":" + port))
}
