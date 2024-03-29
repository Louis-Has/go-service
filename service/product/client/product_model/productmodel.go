// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package product_model

import (
	"context"

	"go-service/internal/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty       = __.Empty
	Product     = __.Product
	ViewProduct = __.ViewProduct

	ProductModel interface {
		GetProductModel(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ViewProduct, error)
	}

	defaultProductModel struct {
		cli zrpc.Client
	}
)

func NewProductModel(cli zrpc.Client) ProductModel {
	return &defaultProductModel{
		cli: cli,
	}
}

func (m *defaultProductModel) GetProductModel(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ViewProduct, error) {
	client := __.NewProductModelClient(m.cli.Conn())
	return client.GetProductModel(ctx, in, opts...)
}
