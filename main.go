package main

import (
	"github.com/gin-gonic/gin"
	"go-service/Controllers"
	"go-service/Models"
)

func init() {
	Models.InitMysql()
}

func main() {
	router := gin.Default()

	Controllers.CoreControl(router)

	if err := router.Run(); err != nil {
		return
	}
}
