package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

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

	findOne, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	_ = copier.Copy(findOne, req)
	err = l.svcCtx.ArticleModel.Update(l.ctx, findOne)
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	_ = copier.Copy(result, findOne)
	return result, nil
}
