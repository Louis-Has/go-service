package AuthorMes

import (
	"github.com/gin-gonic/gin"
	"go-service/Common"
	"go-service/Models"
)

type Service struct {
	Common.BaseService
}

func (receiver Service) GetAll(ctx *gin.Context) {
	var result []Models.AuthorMes

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)
}

func (receiver Service) Inset(ctx *gin.Context) {
	var result Models.AuthorMes

	if err := ctx.ShouldBind(&result); err != nil {
		receiver.Fail(ctx, err)
	}

	if err := Models.Db.Model(&Models.AuthorMes{}).Create(&result); err != nil {
		receiver.Fail(ctx, err)
	}

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)

}
