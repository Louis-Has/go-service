package models

import (
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/redis/go-redis/v9"
	"go-service/common/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var Db *gorm.DB

var Rdb *redis.Client

func init() {
	initMysql()
	initRedis()
}

func initMysql() {
	// db
	var DBError error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Conf.MYSQL.Username,
		conf.Conf.MYSQL.Password,
		conf.Conf.MYSQL.Host,
		conf.Conf.MYSQL.Port,
		conf.Conf.MYSQL.Database,
	)

	Db, DBError = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if DBError != nil {
		log.Printf(color.InRed("db error!"), color.InRed(DBError))
	} else {
		log.Println(color.InGreen(fmt.Sprintf("%s connect Success", conf.Conf.MYSQL.Name)))
	}

	if err := Db.AutoMigrate(&Article{}, &AuthorMes{}); err != nil {
		log.Println(color.InRed(err))
		return
	}

}

func initRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Conf.Redis.Host, conf.Conf.Redis.Port),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(color.InRed(fmt.Sprintf("Redis connect Failed:%s", err)))
		panic(err)
	} else {
		log.Println(color.InGreen(fmt.Sprintf("%s connect Success", conf.Conf.Redis.Name)))
	}
}
