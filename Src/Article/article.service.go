package Article

import (
	"github.com/gin-gonic/gin"
	"go-service/Common"
	"go-service/Models"
)

type Service struct {
	Common.BaseService
}

func (receiver Service) GetAll(ctx *gin.Context) {
	var result []Models.Article
	Models.Db.Model(&Models.Article{}).Find(&result)

	receiver.Success(ctx, result)
}

func (receiver Service) Inset(ctx *gin.Context) {
	result := Models.Article{}

	if err := ctx.ShouldBind(&result); err == nil {

		//fmt.Println(color.InBlue(fmt.Sprintf("%v", result)), err)

		if res := Models.Db.Model(&Models.Article{}).Create(&result); res.Error == nil {

			receiver.Success(ctx, result.ID)
		} else {

			receiver.Fail(ctx, res.Error)
		}
	} else {

		receiver.Fail(ctx, err)
	}
}

func (receiver Service) GetTmp(ctx *gin.Context) {
	var result []Models.Article

	Models.Db.Model(&Models.Article{}).Offset(2).Limit(5).Find(&result)
	receiver.Success(ctx, result)
}

func (receiver Service) GroupBy(ctx *gin.Context) {
	type Group struct {
		Author string `json:"author"`
		Total  int    `json:"total"`
	}

	var result []Group

	Models.Db.Model(&Models.Article{}).Select("author , count(*) as total").Group("author").Find(&result)

	receiver.Success(ctx, result)

}
