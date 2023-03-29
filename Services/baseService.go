package Services

import "github.com/gin-gonic/gin"

type BaseService struct{}

func (con BaseService) Success(context *gin.Context, data interface{}) {

	context.JSON(200, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}

func (con BaseService) Fail(context *gin.Context, err interface{}) {

	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
		"error":   err,
	})
}
