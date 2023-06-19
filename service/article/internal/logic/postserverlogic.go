package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/internal/model"

	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostServerLogic {
	return &PostServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostServerLogic) PostServer(in *art.Article) (*art.ArticleRes, error) {
	data := &model.Article{}
	_ = copier.Copy(data, in)
	insert, err := l.svcCtx.ArticleModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	insertId, _ := insert.LastInsertId()
	findOne, err := l.svcCtx.ArticleModel.FindOne(l.ctx, insertId)
	if err != nil {
		return nil, err
	}

	result := &art.ArticleRes{}
	_ = copier.Copy(result, findOne)

	return result, nil
}
