package usermesmodellogic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/internal/model"

	"go-service/internal/pb"
	"go-service/service/user_mes/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUserMesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostUserMesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUserMesLogic {
	return &PostUserMesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostUserMesLogic) PostUserMes(in *__.User) (*__.User, error) {
	insert, err := l.svcCtx.UserMesModel.Insert(l.ctx, &model.UserMes{
		UserName:     in.UserName,
		SignedPerson: in.SignedPerson,
	})
	if err != nil {
		return nil, err
	}

	LastInsertId, _ := insert.LastInsertId()
	one, err := l.svcCtx.UserMesModel.FindOne(l.ctx, LastInsertId)
	if err != nil {
		return nil, err
	}

	res := &__.User{}
	err = copier.Copy(res, one)
	if err != nil {
		return nil, err
	}

	return res, nil
}
