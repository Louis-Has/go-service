package src

import (
	"github.com/gin-gonic/gin"
	"go-service/src/article"
	"go-service/src/authorMes"
)

func CoreControl(r *gin.Engine) {

	article.Control(r)
	authorMes.Control(r)

}
