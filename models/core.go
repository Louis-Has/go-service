package models

import (
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Config struct {
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

var config Config

var Db *gorm.DB

var Rdb *redis.Client

func InitMysql() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil { // configure config
		log.Println(color.InRed(err))
		return
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Println(color.InRed(err))
		return
	}

	//log.Println(color.InGreen("config load Success"), color.InCyan(config))

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
		log.Printf(color.InRed("db error!"), color.InRed(DBError))
	} else {
		log.Println(color.InGreen(fmt.Sprintf("%s connect Success", config.MYSQL.Name)))
	}

	if err := Db.AutoMigrate(&Article{}, &AuthorMes{}); err != nil {
		log.Println(color.InRed(err))
		return
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Redis.Host, config.Redis.Port),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(color.InRed(fmt.Sprintf("Connect Failed:%s", err)))
		panic(err)
	} else {
		log.Println(color.InGreen(fmt.Sprintf("%s connect Success", config.Redis.Name)))
	}
}
