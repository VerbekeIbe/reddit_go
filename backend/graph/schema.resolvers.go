package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"
	"github.com/google/uuid"

	"github.com/verbekeibe/reddit-backend/graph/generated"
	"github.com/verbekeibe/reddit-backend/graph/model"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.NewCommunity) (*model.Community, error) {
	community:= &model.Community{
		ID: uuid.New().String(),
		Name: input.Name,
		Description: &input.Description,
		DescriptionHTML: &input.Description,
	}
	r.communities = append(r.communities, community)
	return community, nil
}

func (r *queryResolver) Communities(ctx context.Context) ([]*model.Community, error) {
	return r.communities, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

