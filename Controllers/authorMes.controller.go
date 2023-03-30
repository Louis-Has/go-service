package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-service/Services"
)

func AuthorMesControl(r *gin.Engine) {
	rGroup := r.Group("/author")

	{
		rGroup.GET("", Services.AuthorMscService{}.GetAll)
		rGroup.POST("", Services.AuthorMscService{}.Inset)
	}
}
