package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port int
	Host string
}

var config Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // configure config
	if err != nil {
		fmt.Println(err)
		return
	}

	MarRrr := viper.Unmarshal(&config)
	if MarRrr != nil {
		return
	}

	fmt.Println("config load success", config)
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"data":    "mes",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
