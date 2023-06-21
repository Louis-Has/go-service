package article

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutArticleLogic {
	return &PutArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutArticleLogic) PutArticle(req *types.ArticleRes) (resp *types.ArticleRes, err error) {
	tmp := &art.ArticleRes{}
	_ = copier.Copy(tmp, req)

	article, err := l.svcCtx.ArticleClient.PutServer(l.ctx, tmp)
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	_ = copier.Copy(result, article)
	return result, nil
}
