package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlePostLogic {
	return &ArticlePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlePostLogic) ArticlePost(req *types.Article) (resp *types.ArticleRes, err error) {

	tmp := &art.Article{}
	err = copier.Copy(tmp, req)
	if err != nil {
		return nil, err
	}

	article, err := l.svcCtx.ArtServer.PostServer(l.ctx, tmp)
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
