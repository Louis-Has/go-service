package AuthorMes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-service/Common"
	"go-service/Models"
)

type Service struct {
	Common.BaseService
}

type AuthorMes struct {
	Type    string `form:"type" json:"type"`
	Content string `form:"content" json:"content"`
	Author  string `form:"author" json:"author" gorm:"default"`
}

func (receiver Service) GetAll(ctx *gin.Context) {
	var result []Models.AuthorMes

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)
}

func (receiver Service) Inset(ctx *gin.Context) {
	var result Models.AuthorMes

	if err := ctx.ShouldBind(&result); err != nil {
		receiver.Fail(ctx, err)
	}

	if err := Models.Db.Model(&Models.AuthorMes{}).Create(&result).Error; err != nil {
		receiver.Fail(ctx, err)
	}

	Models.Db.Model(&Models.AuthorMes{}).Find(&result)

	receiver.Success(ctx, result)

}

func (receiver Service) TestRedis(ctx *gin.Context) {
	if err := Models.Rdb.Set(ctx, ctx.PostForm("key"), ctx.PostForm("value"), 0).Err(); err != nil {
		panic(err)
	}

	if val, err := Models.Rdb.Get(ctx, ctx.PostForm("key")).Result(); err == redis.Nil {
		receiver.Fail(ctx, "get fail")
	} else {
		receiver.Success(ctx, val)
	}
}

func (receiver Service) TestList(ctx *gin.Context) {
	if _, err := Models.Rdb.LPush(ctx, "listData", ctx.PostForm("value")).Result(); err != nil {
		receiver.Fail(ctx, err)
	} else {
		var result, _ = Models.Rdb.LRange(ctx, "listData", 0, -1).Result()
		receiver.Success(ctx, result)
	}
}
