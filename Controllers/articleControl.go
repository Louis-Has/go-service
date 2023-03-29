package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-service/Services"
)

func ControlRouter(r *gin.Engine) {

	rGroup := r.Group("/work")
	{
		rGroup.GET("", Services.ArticleService{}.Get)
		rGroup.POST("", Services.ArticleService{}.Inset)
	}
}
