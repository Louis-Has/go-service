package AuthorMes

import (
	"github.com/gin-gonic/gin"
)

func Control(r *gin.Engine) {

	rGroup := r.Group("/author")
	{
		rGroup.GET("", GetAll)
		rGroup.POST("", Inset)
		rGroup.POST("/redis", TestRedis)
		rGroup.POST("/redisList", TestList)
	}
}
