package AuthorMes

import (
	"github.com/gin-gonic/gin"
)

func Control(r *gin.Engine) {
	rGroup := r.Group("/author")

	{
		rGroup.GET("", Service{}.GetAll)
		rGroup.POST("", Service{}.Inset)
		rGroup.POST("/redis", Service{}.TestRedis)
		rGroup.POST("/redisList", Service{}.TestList)
	}
}
