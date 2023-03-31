package AuthorMes

import (
	"github.com/gin-gonic/gin"
	"go-service/Common"
	"go-service/Models"
)

type Service struct {
	Common.BaseService
}

type AuthorMes struct {
	Type    string `form:"type" json:"type"`
	Content string `form:"content" json:"content"`
	Author  string `form:"author" json:"author" gorm:"default:author"`
}

func (receiver Service) GetAll(ctx *gin.Context) {
	var result []Models.AuthorMes

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)
}

func (receiver Service) Inset(ctx *gin.Context) {
	var result Models.AuthorMes

	if err := ctx.ShouldBind(&result).Error; err != nil {
		receiver.Fail(ctx, err)
	}

	if err := Models.Db.Model(&Models.AuthorMes{}).Create(&result).Error; err != nil {
		receiver.Fail(ctx, err)
	}

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)

}
