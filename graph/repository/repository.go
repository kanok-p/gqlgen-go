package repository

import (
	"context"

	"graphql-gen/graph/model"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, input *model.CreateInput) (*model.Response, error)
	Update(ctx context.Context, input *model.UpdateInput) (*model.Response, error)
	Get(ctx context.Context, ID string) (*model.Response, error)
	List(ctx context.Context, offset, limit int) (int, []*model.Response, error)
	Delete(ctx context.Context, id string) (*model.Response, error)
}
