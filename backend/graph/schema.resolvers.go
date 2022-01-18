package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/verbekeibe/reddit-backend/graph/generated"
	"github.com/verbekeibe/reddit-backend/graph/model"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.NewCommunity) (*model.Community, error) {
	community := &model.Community{
		ID:              uuid.New().String(),
		Name:            input.Name,
		Description:     &input.Description,
		DescriptionHTML: &input.Description,
	}
	r.communities = append(r.communities, community)
	return community, nil
}

func (r *queryResolver) AllCommunities(ctx context.Context) ([]*model.Community, error) {
	return r.communities, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostsByCommunity(ctx context.Context, communityID string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostsForUser(ctx context.Context, userID string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CommentsByPost(ctx context.Context, postID string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


