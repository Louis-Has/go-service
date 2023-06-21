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

func (c *customArticleModel) checkDeleted(ctx context.Context, id int64) (*Article, error) {
	findOne, err := c.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if findOne.DeletedAt.Valid == true {
		return nil, myError.LoggedOutErr
	}

	return findOne, nil
}

func (c *customArticleModel) SoftFindOne(ctx context.Context, id int64) (*Article, error) {
	checkDeleted, err := c.checkDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	return checkDeleted, nil
}

func (c *customArticleModel) SoftUpdate(ctx context.Context, data *Article) (*Article, error) {
	// check deleted state
	checkDeleted, err := c.checkDeleted(ctx, data.Id)
	if err != nil {
		return nil, err
	}

	// implement incremental updates
	err = copier.Copy(checkDeleted, data)
	if err != nil {
		return nil, err
	}

	// update
	err = c.Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return checkDeleted, nil
}

func (c *customArticleModel) SoftDelete(ctx context.Context, id int64) error {
	_, err := c.checkDeleted(ctx, id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("update %s set deleted_at = ? where `id` = ?", c.table)
	_, err = c.conn.ExecCtx(ctx, query, time.Now(), id)
	return err
}

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn),
	}
}
