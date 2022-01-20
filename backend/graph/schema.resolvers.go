package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/verbekeibe/reddit-backend/graph/generated"
	"github.com/verbekeibe/reddit-backend/graph/model"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.NewCommunity) (*model.Community, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JoinCommunity(ctx context.Context, input model.NewUserCommunity) (*model.UserCommunity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, commentID string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllCommunities(ctx context.Context) ([]*model.Community, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CommunitiesForUser(ctx context.Context, userID string) ([]*model.Community, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByID(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostsByCommunity(ctx context.Context, communityID string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostsForUser(ctx context.Context, userID string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostByID(ctx context.Context, postID string) (*model.Post, error) {
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
