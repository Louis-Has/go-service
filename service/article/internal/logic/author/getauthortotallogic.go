package authorlogic

import (
	"context"
	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorTotalLogic {
	return &GetAuthorTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorTotalLogic) GetAuthorTotal(in *art.Empty) (*art.TotalRes, error) {

	total, err := l.svcCtx.AuthorMesModel.GetArticleTotal(l.ctx, true)
	if err != nil {
		return nil, err
	}

	res := art.TotalRes{AuthorTotals: total}
	return &res, nil
}
