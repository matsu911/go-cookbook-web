package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "public/assets")
	router.LoadHTMLGlob("views/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		content := template.HTML(blackfriday.MarkdownBasic(([]byte)("# test")))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "test",
			"content": content,
		})
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
