package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matsu911/go-cookbook-web/app"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/russross/blackfriday"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.ConnectDB()
	router := gin.Default()
	router.Static("/assets", "public/assets")
	router.LoadHTMLGlob("views/**/*.html")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "test",
		})
	})
	router.GET("/documents/:id", func(c *gin.Context) {
		content := template.HTML(blackfriday.MarkdownBasic(([]byte)("# test")))
		c.HTML(http.StatusOK, "documents/show.html", gin.H{
			"title":   "test",
			"content": content,
		})
	})
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.GET("/", func(c *gin.Context) {
	})
	authorized.GET("/documents/:id", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		c.JSON(http.StatusOK, gin.H{"user": "aa", "secret": "bbb"})
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
