package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlePutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlePutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlePutLogic {
	return &ArticlePutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlePutLogic) ArticlePut(req *types.ArticleRes) (resp *types.ArticleRes, err error) {

	tmp := &art.ArticleRes{}
	_ = copier.Copy(tmp, req)

	article, err := l.svcCtx.ArtServer.PutServer(l.ctx, tmp)
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	_ = copier.Copy(result, article)
	return result, nil
}
