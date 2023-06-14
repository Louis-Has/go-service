package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
