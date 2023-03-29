package Services

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"go-service/Models"
)

type ArticleService struct {
	BaseService
}

func (con ArticleService) Get(context *gin.Context) {
	var articleList []Models.Article
	Models.Db.Find(&articleList)

	con.Success(context, articleList)
}

func (con ArticleService) Inset(context *gin.Context) {
	articleData := Models.Article{}

	if err := context.ShouldBind(&articleData); err == nil {

		fmt.Println(color.InBlue(fmt.Sprintf("%v", articleData)), err)

		if result := Models.Db.Create(&articleData); result.Error == nil {

			con.Success(context, articleData.ID)
		} else {

			con.Fail(context, result.Error)
		}
	} else {

		con.Fail(context, err)
	}

}
