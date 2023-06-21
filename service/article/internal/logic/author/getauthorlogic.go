package authorlogic

import (
	"context"
	"github.com/jinzhu/copier"

	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorLogic {
	return &GetAuthorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorLogic) GetAuthor(in *art.Id) (*art.AuthorRes, error) {

	authors, err := l.svcCtx.AuthorMesModel.FindAuthorsById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	result := &art.AuthorRes{}
	err = copier.Copy(&result.Authors, &authors)
	if err != nil {
		return nil, err
	}

	return result, nil
}
