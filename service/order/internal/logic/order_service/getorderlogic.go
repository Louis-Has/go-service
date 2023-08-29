package orderservicelogic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"time"

	"go-service/internal/pb"
	"go-service/service/order/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *__.Id) (*__.WholeOrder, error) {
	res := &__.WholeOrder{}

	findOrder, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	res.Id = findOrder.Id
	res.UserId = findOrder.UserId

	findOneReMes, err := l.svcCtx.OrderReceiveMesModel.FindOne(l.ctx, findOrder.OrderReceiveMesId)
	if err != nil {
		return nil, err
	}

	res.ReceiveMes = &__.OrderReceive{}
	err = copier.Copy(res.ReceiveMes, findOneReMes)
	if err != nil {
		return nil, fmt.Errorf("%v coi %w", time.Time{}, err)
	}

	// details
	findAllDetail, err := l.svcCtx.OrderDetailModel.FindAllByOrderId(l.ctx, findOrder.Id)
	if err != nil {
		return nil, err
	}

	// find all product name
	foundAllProduct, err := l.svcCtx.ProductModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	productNameMap := make(map[int64]string)

	for _, v := range foundAllProduct {
		productNameMap[v.Id] = v.ProductName
	}

	for _, detail := range findAllDetail {
		tmp := &__.OrderDetailRes{}
		tmp.ProductName = productNameMap[detail.ProductId]

		err := copier.Copy(tmp, detail)
		if err != nil {
			return nil, err
		}
		res.Details = append(res.Details, tmp)
	}

	return res, nil
}
