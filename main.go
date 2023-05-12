package main

import (
	"github.com/gin-gonic/gin"
	"go-service/models"
	"go-service/src"
)

func init() {
	models.InitMysql()
}

func main() {
	router := gin.Default()

	src.CoreControl(router)

	if err := router.Run(":8080"); err != nil {
		return
	}
}
