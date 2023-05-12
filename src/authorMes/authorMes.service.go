package authorMes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-service/common"
	"go-service/models"
)

type AuthorMes struct {
	Type    string `form:"type" json:"type"`
	Content string `form:"content" json:"content"`
	Author  string `form:"author" json:"author" gorm:"default"`
}

func GetAll(ctx *gin.Context) {
	var result []models.AuthorMes

	models.Db.Model(&models.AuthorMes{}).Find(&result)

	common.Success(ctx, result)
}

func Inset(ctx *gin.Context) {
	var result models.AuthorMes

	if err := ctx.ShouldBindJSON(&result); err != nil {
		common.Fail(ctx, err)
	}

	if err := models.Db.Model(&models.AuthorMes{}).Create(&result).Error; err != nil {
		common.Fail(ctx, err)
	}

	models.Db.Model(&models.AuthorMes{}).Find(&result)

	common.Success(ctx, result)

}

func TestRedis(ctx *gin.Context) {
	if err := models.Rdb.Set(ctx, ctx.PostForm("key"), ctx.PostForm("value"), 0).Err(); err != nil {
		panic(err)
	}

	if val, err := models.Rdb.Get(ctx, ctx.PostForm("key")).Result(); err == redis.Nil {
		common.Fail(ctx, "get fail")
	} else {
		common.Success(ctx, val)
	}
}

func TestList(ctx *gin.Context) {
	if _, err := models.Rdb.LPush(ctx, "listData", ctx.PostForm("value")).Result(); err != nil {
		common.Fail(ctx, err)
	} else {
		var result, _ = models.Rdb.LRange(ctx, "listData", 0, -1).Result()
		common.Success(ctx, result)
	}
}
