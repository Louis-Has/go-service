package Common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
	return
}

func Fail(ctx *gin.Context, err interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"error":   err,
	})
	return
}
