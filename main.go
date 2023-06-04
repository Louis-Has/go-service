package main

import (
	"github.com/gin-gonic/gin"
	"go-service/src"
)

func init() {}

func main() {
	router := gin.Default()
	src.CoreControl(router)

	if err := router.Run(":8080"); err != nil {
		return
	}
}
