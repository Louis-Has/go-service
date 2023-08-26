package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserMesModel = (*customUserMesModel)(nil)

type (
	// UserMesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMesModel.
	UserMesModel interface {
		userMesModel
	}

	customUserMesModel struct {
		*defaultUserMesModel
	}
)

// NewUserMesModel returns a model for the database table.
func NewUserMesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserMesModel {
	return &customUserMesModel{
		defaultUserMesModel: newUserMesModel(conn, c, opts...),
	}
}
