package article

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.PathID) (resp *types.ArticleRes, err error) {
	article, err := l.svcCtx.ArticleClient.GetServer(l.ctx, &art.Id{Id: req.Id})
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
