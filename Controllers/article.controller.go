package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-service/Services"
)

func ArticleControl(r *gin.Engine) {

	rGroup := r.Group("/article")

	{
		rGroup.GET("", Services.ArticleService{}.GetAll)
		rGroup.POST("", Services.ArticleService{}.Inset)
		rGroup.GET("/limit", Services.ArticleService{}.GetTmp)
		rGroup.GET("/group", Services.ArticleService{}.GroupBy)
	}
}
