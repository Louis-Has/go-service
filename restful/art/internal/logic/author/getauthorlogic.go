package author

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorLogic {
	return &GetAuthorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorLogic) GetAuthor(req *types.PathID) (resp *types.AuthorRes, err error) {
	author, err := l.svcCtx.AuthorClient.GetAuthor(l.ctx, &art.Id{Id: req.Id})
	if err != nil {
		return nil, err
	}
	result := &types.AuthorRes{}
	err = copier.Copy(&result, &author)
	if err != nil {
		return nil, err
	}

	return result, nil
}
