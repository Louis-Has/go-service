package Article

import (
	"github.com/gin-gonic/gin"
	"go-service/Common"
	"go-service/Models"
)

func GetAll(ctx *gin.Context) {
	var result []Models.Article
	Models.Db.Model(&Models.Article{}).Find(&result)

	Common.Success(ctx, result)
}

func Inset(ctx *gin.Context) {
	result := Models.Article{}

	if err := ctx.ShouldBind(&result); err == nil {

		//fmt.Println(color.InBlue(fmt.Sprintf("%v", result)), err)

		if res := Models.Db.Model(&Models.Article{}).Create(&result); res.Error == nil {

			Common.Success(ctx, result.ID)
		} else {

			Common.Fail(ctx, res.Error)
		}
	} else {

		Common.Fail(ctx, err)
	}
}

func GetTmp(ctx *gin.Context) {
	var result []Models.Article

	Models.Db.Model(&Models.Article{}).Offset(2).Limit(5).Find(&result)
	Common.Success(ctx, result)
}

func GroupBy(ctx *gin.Context) {
	type Group struct {
		Author string `json:"author"`
		Total  int    `json:"total"`
	}

	var result []Group

	Models.Db.Model(&Models.Article{}).Select("author , count(*) as total").Group("author").Find(&result)

	Common.Success(ctx, result)

}
