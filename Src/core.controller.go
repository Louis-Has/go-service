package Src

import (
	"github.com/gin-gonic/gin"
	"go-service/Src/article"
	"go-service/Src/authorMes"
)

func CoreControl(r *gin.Engine) {

	article.Control(r)
	authorMes.Control(r)

}
