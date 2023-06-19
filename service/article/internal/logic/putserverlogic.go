package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-service/internal/model"

	"go-service/service/article/internal/svc"
	"go-service/service/pb/art"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutServerLogic {
	return &PutServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutServerLogic) PutServer(in *art.ArticleRes) (*art.ArticleRes, error) {

	tmp := &model.Article{}
	err := copier.Copy(tmp, in)
	if err != nil {
		return nil, err
	}

	update, err := l.svcCtx.ArticleModel.SoftUpdate(l.ctx, tmp)
	if err != nil {
		return nil, err
	}

	result := &art.ArticleRes{}
	_ = copier.Copy(result, update)
	fmt.Printf("this is result :%+v\n", result)
	return result, nil
}
