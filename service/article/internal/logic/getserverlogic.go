package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerLogic {
	return &GetServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServerLogic) GetServer(in *art.ArticleId) (*art.ArticleRes, error) {
	findOne, err := l.svcCtx.ArticleModel.SoftFindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	result := &art.ArticleRes{}
	err = copier.Copy(&result, findOne)
	if err != nil {
		return nil, err
	}

	return result, nil
}
