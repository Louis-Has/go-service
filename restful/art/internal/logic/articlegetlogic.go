package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

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

func (l *ArticleGetLogic) ArticleGet(req *types.ArticleId) (resp *types.ArticleRes, err error) {
	findOne, err := l.svcCtx.ArticleModel.SoftFindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	_ = copier.Copy(&result, findOne)

	return result, nil
}
