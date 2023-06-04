package conf

import (
	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
	"log"
)

type confType struct {
	Port int
	Host string
	MYSQL
	Redis
}

type MYSQL struct {
	Name, Host, Database, Username, Password string
	Port                                     int
}

type Redis struct {
	Host, Name string
	Port       int
}

var Conf confType

func init() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil { // configure config
		log.Println(color.InRed(err))
		return
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Println(color.InRed(err))
		return
	}

	//log.Println(color.InGreen("config load Success"), color.InCyan(config))
}
