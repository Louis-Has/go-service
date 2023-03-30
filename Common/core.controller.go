package Common

import (
	"github.com/gin-gonic/gin"
	"go-service/Src/Article"
	"go-service/Src/AuthorMes"
)

func CoreControl(r *gin.Engine) {

	Article.Control(r)
	AuthorMes.Control(r)

}
