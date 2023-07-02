package authorlogic

import (
	"context"
	"fmt"

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

func (l *GetAuthorTotalLogic) GetAuthorTotal(in *art.NeedLived) (*art.TotalRes, error) {
	_ = in

	fmt.Printf("this is ready\n")
	total, err := l.svcCtx.AuthorMesModel.GetArticleTotal(l.ctx, in.Lived)
	fmt.Printf("this is the%+v\n", total)
	if err != nil {
		return nil, err
	}

	res := art.TotalRes{AuthorTotals: total}
	fmt.Printf("this is rest res:%+v\n", res)

	return &res, nil
}
