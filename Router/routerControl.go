package Router

import (
	"github.com/gin-gonic/gin"
)

func ControlRouter(r *gin.Engine) {

	rGroup := r.Group("/group")
	{
		rGroup.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"success": true,
				"data":    "mes",
			})
		})
	}
}
