package orderservicelogic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go-service/internal/pb"
	"go-service/service/order/internal/svc"
)

type GetUserOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrderLogic {
	return &GetUserOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserOrderLogic) GetUserOrder(in *__.Id) (*__.UserOrderAll, error) {
	res := &__.UserOrderAll{}

	foundUserMes, err := l.svcCtx.UserMesModelClient.GetUserMes(l.ctx, &__.ReqID{Id: in.Id})
	if err != nil {
		return nil, err
	}

	// find user_mes
	res.Id = in.Id
	res.UserName = foundUserMes.UserName

	// find orders
	orders, err := l.svcCtx.OrderModel.FindOrderIds(l.ctx, in.Id)
	if err != nil {
		return nil, fmt.Errorf("orders %w", err)
	}

	for _, order := range orders {

		resOrder := &__.WholeOrder{}
		err = copier.Copy(resOrder, order)
		if err != nil {
			return nil, fmt.Errorf("aaa %w", err)
		}

		// find reMes
		foundReMes, err := l.svcCtx.OrderReceiveMesModel.FindOne(l.ctx, order.OrderReceiveMesId)
		if err != nil {
			return nil, fmt.Errorf("foundReMes %w", err)
		}

		resRe := &__.OrderReceive{}
		err = copier.Copy(resRe, foundReMes)
		if err != nil {
			return nil, fmt.Errorf("resRe %w", err)
		}
		resOrder.ReceiveMes = resRe

		// find details
		foundAllDetails, err := l.svcCtx.OrderDetailModel.FindAllByOrderId(l.ctx, order.Id)
		if err != nil {
			return nil, fmt.Errorf("foundAllDetails %w", err)
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

		for _, foundDetail := range foundAllDetails {
			resDetail := &__.OrderDetailRes{}
			resDetail.ProductName = productNameMap[foundDetail.ProductId]

			err := copier.Copy(resDetail, foundDetail)
			if err != nil {
				return nil, fmt.Errorf("resDetail %w", err)
			}

			resOrder.Details = append(resOrder.Details, resDetail)
		}

		res.Orders = append(res.Orders, resOrder)
	}

	return res, nil

}
