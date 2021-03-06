package repository

import (
	"context"

	"github.com/monmaru/gae-graphql/domain/model"
)

type BlogRepository interface {
	Create(ctx context.Context, blog *model.Blog) (*model.Blog, error)
	CreateMulti(ctx context.Context, blogs []*model.Blog) ([]*model.Blog, error)
	NewQuery() Query
}

type Query interface {
	Limit(limit int) Query
	Offset(offset int) Query
	Order(filedName string) Query
	Filter(filterStr string, value interface{}) Query
	GetAll(ctx context.Context) (*model.BlogList, error)
}
