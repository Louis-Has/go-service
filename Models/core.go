package Models

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func InitMysql() {
	viper.AddConfigPath("./Conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil { // configure config
		fmt.Println(color.InRed(err))
		return
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(color.InRed(err))
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

	Db, DBError = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if DBError != nil {
		fmt.Printf(color.InRed("db error!"), color.InRed(DBError))
	} else {
		fmt.Println(color.InGreen(fmt.Sprintf("%s connect Success", config.MYSQL.Name)))
	}

	if err := Db.AutoMigrate(&Article{}, &AuthorMes{}); err != nil {
		fmt.Println(color.InRed(err))
		return
	}

}
