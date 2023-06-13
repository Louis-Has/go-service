package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleLogic) Article(req *types.GetArticleReq) (res *types.Article, err error) {

	findOne, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var findResult types.Article
	_ = copier.Copy(&findResult, findOne)

	return &findResult, nil
}
