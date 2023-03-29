package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Type    string `form:"type" json:"type"`
	Content string `form:"content" json:"content"`
	Author  string `form:"author" json:"author"`
}
