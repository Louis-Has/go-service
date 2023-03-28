package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-service/Router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Port int
	Host string
	MYSQL
}

type MYSQL struct {
	Name, Host, Database, Username, Password string
	Port                                     int
}

var config Config

var Db *gorm.DB

func init() {
	viper.AddConfigPath("./Conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil { // configure config
		fmt.Println(color.InRed(err))
		return
	}

	if MarErr := viper.Unmarshal(&config); MarErr != nil {
		return
	}

	fmt.Println(color.InGreen("config load Success"), color.InCyan(config))

	// db
	var DBError error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MYSQL.Username,
		config.MYSQL.Password,
		config.MYSQL.Host,
		config.MYSQL.Port,
		config.MYSQL.Database,
	)
	_, DBError = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if DBError != nil {
		fmt.Printf(color.InRed("db error!"), color.InRed(DBError))
	} else {
		fmt.Println(color.InGreen(fmt.Sprintf("%s connect Success", config.MYSQL.Name)))
	}

}

func main() {
	router := gin.Default()

	Router.ControlRouter(router) //mount Router

	err := router.Run()
	if err != nil {
		return
	}
}
