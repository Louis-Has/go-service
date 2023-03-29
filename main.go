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
	Controllers.ControlRouter(router) //mount Router

	if err := router.Run(); err != nil {
		return
	}
}
