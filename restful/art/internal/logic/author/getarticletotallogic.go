package author

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleTotalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleTotalLogic {
	return &GetArticleTotalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleTotalLogic) GetArticleTotal(req *types.NeedLived) (resp *types.TotalRes, err error) {
	total, err := l.svcCtx.AuthorClient.GetAuthorTotal(l.ctx, &art.NeedLived{Lived: req.Lived})
	if err != nil {
		return nil, err
	}

	res := &types.TotalRes{}
	err = copier.Copy(&res, total)
	if err != nil {
		return nil, err
	}

	return res, nil
}
