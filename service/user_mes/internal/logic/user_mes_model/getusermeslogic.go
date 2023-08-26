package usermesmodellogic

import (
	"context"
	"github.com/jinzhu/copier"
	myError "go-service/internal/error"

	"go-service/internal/pb"
	"go-service/service/user_mes/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMesLogic {
	return &GetUserMesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserMesLogic) GetUserMes(in *__.ReqID) (*__.User, error) {
	one, err := l.svcCtx.UserMesModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, myError.UserNotExist
	}

	result := &__.User{}
	err = copier.Copy(result, one)
	if err != nil {
		return nil, err
	}

	return result, nil
}
