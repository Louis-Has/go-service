package orderservicelogic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	myError "go-service/internal/error"
	"go-service/internal/model"
	"go-service/internal/pb"
	"go-service/service/order/internal/svc"
	"time"
)

type PostOrderMesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostOrderMesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostOrderMesLogic {
	return &PostOrderMesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostOrderMesLogic) PostOrderMes(in *__.OrderReq) (*__.WholeOrder, error) {
	if len(in.Details) == 0 {
		return nil, myError.OrderDetailNotExist
	}

	insertOrder, err := l.svcCtx.OrderModel.Insert(l.ctx, &model.Order{
		UserId:            in.UserId,
		OrderReceiveMesId: in.OrderReceiveMesId,
		Payment:           float64(in.Payment),
		PaymentType:       in.PaymentType,
		Postage:           in.Postage,
	})
	if err != nil {
		return nil, err
	}

	insertId, _ := insertOrder.LastInsertId()

	for _, order := range in.Details {
		_, err := l.svcCtx.OrderDetailModel.Insert(l.ctx, &model.OrderDetail{
			OrderId:          insertId,
			UserId:           in.UserId,
			ProductId:        order.ProductId,
			CurrentUnitPrice: float64(order.CurrentUnitPrice),
			ProductQuantity:  order.ProductQuantity,
			TotalPrice:       float64(order.CurrentUnitPrice) * float64(order.ProductQuantity),
		})
		if err != nil {
			return nil, err
		}
	}

	lastId, _ := insertOrder.LastInsertId()
	res := &__.WholeOrder{}

	findOrder, err := l.svcCtx.OrderModel.FindOne(l.ctx, lastId)
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
