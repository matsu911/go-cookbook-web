package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "public/assets")
	router.LoadHTMLGlob("views/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "test",
		})
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
