package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"graphql-gen/graph/generated"
	"graphql-gen/graph/model"
)

func (r *mutationResolver) Create(ctx context.Context, input *model.CreateInput) (*model.Response, error) {
	return r.Repo.Create(ctx, input)
}

func (r *mutationResolver) Update(ctx context.Context, input *model.UpdateInput) (*model.Response, error) {
	return r.Repo.Update(ctx, input)
}

func (r *mutationResolver) Delete(ctx context.Context, input string) (*model.Response, error) {
	return r.Repo.Delete(ctx, input)
}

func (r *queryResolver) List(ctx context.Context, input *model.PaginationInput) (*model.ResponseList, error) {
	var limit, offset int
	resp := &model.ResponseList{}
	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Offset != nil {
			offset = *input.Offset
		}
	}
	total, list, err := r.Repo.List(ctx, offset, limit)
	if err != nil {
		return resp, err
	}
	resp.Total = &total
	resp.Response = list

	return resp, nil
}

func (r *queryResolver) Read(ctx context.Context, input string) (*model.Response, error) {
	return r.Repo.Get(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
