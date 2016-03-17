package models

import (
	"html/template"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/russross/blackfriday"
)

type Document struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Text  string `gorm:"type:text"`
	Slug  string `gorm:"size:255;index"`
}

func (doc Document) Content() interface{} {
	return template.HTML(string(blackfriday.MarkdownBasic(([]byte)(doc.Text))))
	// return template.HTML()
}
