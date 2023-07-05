package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	myError "go-service/internal/error"
	"go-service/service/pb/art"
)

var _ AuthorMesModel = (*customAuthorMesModel)(nil)

type (
	// AuthorMesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthorMesModel.
	AuthorMesModel interface {
		authorMesModel
		checkDeleted(ctx context.Context, id int64) (*AuthorMes, error)
		FindAuthorsById(ctx context.Context, id int64) ([]*AuthorMes, error)
		GetArticleTotal(ctx context.Context, lived bool) ([]*art.AuthorTotal, error)
	}

	customAuthorMesModel struct {
		*defaultAuthorMesModel
	}
)

func (c customAuthorMesModel) checkDeleted(ctx context.Context, id int64) (*AuthorMes, error) {
	findOne, err := c.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if findOne.DeletedAt.Valid == true {
		return nil, myError.LoggedOutErr
	}

	return findOne, nil
}

func (c customAuthorMesModel) FindAuthorsById(ctx context.Context, id int64) ([]*AuthorMes, error) {
	articleTableName := "article"
	query := fmt.Sprintf("select * from %s as t1 join %s as t2 on t1.id = ? and t1.author = t2.author", articleTableName, c.table)
	var resp []*AuthorMes
	err := c.conn.QueryRowsCtx(ctx, &resp, query, id)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customAuthorMesModel) GetArticleTotal(ctx context.Context, lived bool) ([]*art.AuthorTotal, error) {
	_ = lived
	var res []*art.AuthorTotal

	query := fmt.Sprintf("SELECT t1.id ,t1.author ,IF(t1.deleted_at,TRUE,FALSE) deleted ,COUNT(t2.id) mesCount ,IFNULL(SUM(t2.cash),0) cashSum " +
		"from article t1 LEFT JOIN author_mes t2 ON t1.author = t2.author  GROUP BY t1.author")
	err := c.conn.QueryRowsCtx(ctx, &res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewAuthorMesModel returns a model for the database table.
func NewAuthorMesModel(conn sqlx.SqlConn) AuthorMesModel {
	return &customAuthorMesModel{
		defaultAuthorMesModel: newAuthorMesModel(conn),
	}
}
