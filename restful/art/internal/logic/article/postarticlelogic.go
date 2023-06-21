package article

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostArticleLogic {
	return &PostArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostArticleLogic) PostArticle(req *types.Article) (resp *types.ArticleRes, err error) {
	tmp := &art.Article{}
	err = copier.Copy(tmp, req)
	if err != nil {
		return nil, err
	}

	article, err := l.svcCtx.ArticleClient.PostServer(l.ctx, tmp)
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
