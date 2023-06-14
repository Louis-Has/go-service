package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-service/internal/model"
	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlePostLogic {
	return &ArticlePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlePostLogic) ArticlePost(req *types.Article) (resp *types.ArticleRes, err error) {

	data := &model.Article{}
	_ = copier.Copy(data, req)
	insert, err := l.svcCtx.ArticleModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	insertId, _ := insert.LastInsertId()
	findOne, err := l.svcCtx.ArticleModel.FindOne(l.ctx, insertId)
	if err != nil {
		return nil, err
	}

	result := &types.ArticleRes{}
	_ = copier.Copy(result, findOne)

	return result, nil
}
