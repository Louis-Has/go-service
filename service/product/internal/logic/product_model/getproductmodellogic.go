package productmodellogic

import (
	"context"
	"go-service/internal/pb"
	"go-service/service/product/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductModelLogic {
	return &GetProductModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductModelLogic) GetProductModel(in *__.Empty) (*__.ViewProduct, error) {
	findAllCate, err := l.svcCtx.ProductCategoryModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	cateMap := make(map[int64]string)
	for _, v := range findAllCate {
		cateMap[v.Id] = v.Name
	}

	findAllProduct, err := l.svcCtx.ProductModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	res := &__.ViewProduct{}
	for _, v := range findAllProduct {
		res.Products = append(res.Products, &__.Product{
			Category:    cateMap[v.CateId],
			ProductName: v.ProductName,
			Subtitle:    v.Subtitle,
			Price:       float32(v.Price),
			Stock:       v.Stock,
		})

	}

	return res, nil
}
