package article

import (
	"context"
	"go-service/service/pb/art"

	"go-service/restful/art/internal/svc"
	"go-service/restful/art/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.PathID) error {
	_, err := l.svcCtx.ArticleClient.DeleteServer(l.ctx, &art.Id{Id: req.Id})
	if err != nil {
		return err
	}

	return nil
}
