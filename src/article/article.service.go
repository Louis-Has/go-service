package article

import (
	"github.com/gin-gonic/gin"
	"go-service/common"
	"go-service/models"
)

func GetAll(ctx *gin.Context) {
	var result []models.Article
	models.Db.Model(&models.Article{}).Find(&result)

	common.Success(ctx, result)
}

func Inset(ctx *gin.Context) {
	result := models.Article{}

	if err := ctx.ShouldBind(&result); err == nil {

		//fmt.Println(color.InBlue(fmt.Sprintf("%v", result)), err)

		if res := models.Db.Model(&models.Article{}).Create(&result); res.Error == nil {

			common.Success(ctx, result.ID)
		} else {

			common.Fail(ctx, res.Error)
		}
	} else {

		common.Fail(ctx, err)
	}
}

func GetTmp(ctx *gin.Context) {
	var result []models.Article

	models.Db.Model(&models.Article{}).Offset(2).Limit(5).Find(&result)
	common.Success(ctx, result)
}

func GroupBy(ctx *gin.Context) {
	type Group struct {
		Author string `json:"author"`
		Total  int    `json:"total"`
	}

	var result []Group

	models.Db.Model(&models.Article{}).Select("author , count(*) as total").Group("author").Find(&result)

	common.Success(ctx, result)

}
