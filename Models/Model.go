package Models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Type    string `form:"type" json:"type"`
	Content string `form:"content" json:"content"`
	Author  string `form:"author" json:"author" gorm:"default:author"`
}

type AuthorMes struct {
	gorm.Model
	Author       string `form:"author" json:"author" gorm:"default:author;uniqueIndex"`
	SignedPerson bool   `form:"signed_person" json:"signed_person" gorm:"default:false"`
	Cash         int    `form:"cash" json:"cash"`
}
