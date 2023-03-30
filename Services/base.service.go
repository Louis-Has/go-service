package Services

import "github.com/gin-gonic/gin"

type BaseService struct{}

func (con BaseService) Success(ctx *gin.Context, data interface{}) {

	ctx.JSON(200, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}

func (con BaseService) Fail(ctx *gin.Context, err interface{}) {

	ctx.JSON(200, gin.H{
		"success": true,
		"data":    nil,
		"error":   err,
	})
}
