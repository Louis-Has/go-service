package Article

import (
	"github.com/gin-gonic/gin"
)

func Control(r *gin.Engine) {

	rGroup := r.Group("/article")

	{
		rGroup.GET("", GetAll)
		rGroup.POST("", Inset)
		rGroup.GET("/limit", GetTmp)
		rGroup.GET("/group", GroupBy)
	}
}
