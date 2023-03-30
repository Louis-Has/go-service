package Services

import (
	"github.com/gin-gonic/gin"
	"go-service/Models"
)

type AuthorMscService struct {
	BaseService
}

func (receiver AuthorMscService) GetAll(ctx *gin.Context) {
	var result []Models.AuthorMes

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)
}

func (receiver AuthorMscService) Inset(ctx *gin.Context) {
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
