package logic

import (
	"context"

	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServerLogic {
	return &DeleteServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteServerLogic) DeleteServer(in *art.ArticleId) (*art.NilRes, error) {
	err := l.svcCtx.ArticleModel.SoftDelete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &art.NilRes{}, nil
}
