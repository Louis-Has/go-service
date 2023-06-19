package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	myError "go-service/internal/error"
	"time"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		checkDeleted(ctx context.Context, id int64) (*Article, error)
		SoftFindOne(ctx context.Context, id int64) (*Article, error)
		SoftUpdate(ctx context.Context, data *Article) (*Article, error)
		SoftDelete(ctx context.Context, id int64) error
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

func (m *customArticleModel) checkDeleted(ctx context.Context, id int64) (*Article, error) {
	findOne, err := m.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if findOne.DeletedAt.Valid == true {
		return nil, myError.LoggedOutErr
	}

	return findOne, nil
}

func (m *customArticleModel) SoftFindOne(ctx context.Context, id int64) (*Article, error) {
	checkDeleted, err := m.checkDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	return checkDeleted, nil
}

func (m *customArticleModel) SoftUpdate(ctx context.Context, data *Article) (*Article, error) {
	// check deleted state
	checkDeleted, err := m.checkDeleted(ctx, data.Id)
	if err != nil {
		return nil, err
	}

	// implement incremental updates
	err = copier.Copy(checkDeleted, data)
	if err != nil {
		return nil, err
	}

	// update
	err = m.Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return checkDeleted, nil
}

func (m *customArticleModel) SoftDelete(ctx context.Context, id int64) error {
	_, err := m.checkDeleted(ctx, id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("update %s set deleted_at = ? where `id` = ?", m.table)
	_, err = m.conn.ExecCtx(ctx, query, time.Now(), id)
	return err
}

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn),
	}
}
