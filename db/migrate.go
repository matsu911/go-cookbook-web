package main

import (
	"github.com/matsu911/go-cookbook-web/app"
	"github.com/matsu911/go-cookbook-web/app/models"
)

func main() {
	db := app.ConnectDB()
	db.CreateTable(&models.Document{})
}
