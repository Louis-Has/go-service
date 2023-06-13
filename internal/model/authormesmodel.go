package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AuthorMesModel = (*customAuthorMesModel)(nil)

type (
	// AuthorMesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthorMesModel.
	AuthorMesModel interface {
		authorMesModel
	}

	customAuthorMesModel struct {
		*defaultAuthorMesModel
	}
)

// NewAuthorMesModel returns a model for the database table.
func NewAuthorMesModel(conn sqlx.SqlConn) AuthorMesModel {
	return &customAuthorMesModel{
		defaultAuthorMesModel: newAuthorMesModel(conn),
	}
}
