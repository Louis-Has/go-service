package main

import (
	"github.com/gin-gonic/gin"
	"go-service/Models"
	"go-service/Src"
)

func init() {
	Models.InitMysql()
}

func main() {
	router := gin.Default()

	Src.CoreControl(router)

	if err := router.Run(":8080"); err != nil {
		return
	}
}
