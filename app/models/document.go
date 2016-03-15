package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Document struct {
	gorm.Model
	Text string `gorm:"type:text"`
	Slug string `gorm:"size:255;index"`
}
