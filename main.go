package main

import (
	"github.com/gin-gonic/gin"
	"go-service/Common"
	"go-service/Models"
)

func init() {
	Models.InitMysql()
}

func main() {
	router := gin.Default()

	Common.CoreControl(router)

	if err := router.Run(); err != nil {
		return
	}
}
