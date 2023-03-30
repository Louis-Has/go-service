package Article

import (
	"github.com/gin-gonic/gin"
)

func Control(r *gin.Engine) {

	rGroup := r.Group("/article")

	{
		rGroup.GET("", Service{}.GetAll)
		rGroup.POST("", Service{}.Inset)
		rGroup.GET("/limit", Service{}.GetTmp)
		rGroup.GET("/group", Service{}.GroupBy)
	}
}
