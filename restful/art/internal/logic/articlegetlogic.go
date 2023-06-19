package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleGetLogic {
	return &ArticleGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleGetLogic) ArticleGet(req *types.PathID) (resp *types.ArticleRes, err error) {
	article, err := l.svcCtx.ArtServer.GetServer(l.ctx, &art.ArticleId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	err = copier.Copy(result, article)
	if err != nil {
		return nil, err
	}

	return result, nil
}
