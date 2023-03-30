package Controllers

import "github.com/gin-gonic/gin"

func CoreControl(r *gin.Engine) {

	ArticleControl(r)
	AuthorMesControl(r)

}
