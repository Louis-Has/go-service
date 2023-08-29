package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-service/internal/model"
	__ "go-service/internal/pb"
	"go-service/service/order/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.OrderModel
	model.OrderDetailModel
	model.OrderReceiveMesModel
	__.UserMesModelClient
	model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.UserMesServer,
	})
	return &ServiceContext{
		Config:               c,
		OrderModel:           model.NewOrderModel(db, c.CacheRedis),
		OrderDetailModel:     model.NewOrderDetailModel(db, c.CacheRedis),
		OrderReceiveMesModel: model.NewOrderReceiveMesModel(db, c.CacheRedis),
		UserMesModelClient:   __.NewUserMesModelClient(conn.Conn()),
		ProductModel:         model.NewProductModel(db, c.CacheRedis),
	}
}
