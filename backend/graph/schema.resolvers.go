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
	var communities []*model.Community

	r.DB.Find(&communities)

	return communities, nil
}

func (r *queryResolver) CommunitiesForUser(ctx context.Context, userID string) ([]*model.Community, error) {
	var communities []*model.Community
	var userCommunities []*model.UserCommunity
	var communityIDs []string

	r.DB.Where("user_id = ?", userID).Find(&userCommunities)

	for _, community := range userCommunities {
		communityIDs = append(communityIDs, community.CommunityID)
	}
	r.DB.Find(&communities, communityIDs)

	return communities, nil
}

func (r *queryResolver) PostsForUser(ctx context.Context, userID string) ([]*model.Post, error) {
	var userCommunities []*model.UserCommunity
	var communityIDs []string
	var posts []*model.Post

	r.DB.Where("user_id = ?", userID).Find(&userCommunities)

	for _, community := range userCommunities {
		communityIDs = append(communityIDs, community.CommunityID)
	}
	r.DB.Where("community_id IN ?", communityIDs).Find(&posts)

	return posts, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	r.DB.Find(&users)

	return users, nil
}

func (r *queryResolver) UserByID(ctx context.Context, userID string) (*model.User, error) {
	var user *model.User

	r.DB.Where("id = ?", userID).Find(&user)
	return user, nil
}

func (r *queryResolver) PostsByCommunity(ctx context.Context, communityID string) ([]*model.Post, error) {
	var posts []*model.Post

	r.DB.Where("community_id = ?", communityID).Find(&posts)
	return posts, nil
}



func (r *queryResolver) PostByID(ctx context.Context, postID string) (*model.Post, error) {
	var post *model.Post

	r.DB.Where("id = ?", postID).Find(&post)
	return post, nil
}

func (r *queryResolver) CommentsByPost(ctx context.Context, postID string) ([]*model.Comment, error) {
	var comments []*model.Comment

	r.DB.Where("post_id = ?", postID).Find(&comments)
	return comments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
