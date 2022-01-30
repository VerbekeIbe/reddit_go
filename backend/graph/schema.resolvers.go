package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/verbekeibe/reddit-backend/graph/generated"
	"github.com/verbekeibe/reddit-backend/graph/model"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.NewCommunity) (*model.Community, error) {
	var description = input.Description
	var descriptionHTML = input.DescriptionHTML
	community := &model.Community{
		ID:              uuid.New().String(),
		Name:            input.Name,
		Description:     &description,
		DescriptionHTML: &descriptionHTML,
	}

	r.DB.Create(&community)

	return community, nil
}

func (r *mutationResolver) JoinCommunity(ctx context.Context, input model.NewUserCommunity) (*model.UserCommunity, error) {
	userCommunity := &model.UserCommunity{
		UserID:      input.UserID,
		CommunityID: input.CommunityID,
	}
	r.DB.Create(&userCommunity)

	return userCommunity, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	timeNow := time.Now()
	timestamp := int(timeNow.Unix())

	post := &model.Post{
		ID:          uuid.New().String(),
		CommunityID: input.CommunityID,
		UserID:      input.UserID,
		Title:       input.Title,
		Content:     input.Content,
		ContentHTML: input.ContentHTML,
		Timestamp:   timestamp,
	}
	r.DB.Create(&post)

	return post, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	timeNow := time.Now()
	timestamp := int(timeNow.Unix())

	comment := &model.Comment{
		ID:        uuid.New().String(),
		PostID:    input.PostID,
		UserID:    input.UserID,
		Content:   input.Content,
		Timestamp: timestamp,
	}

	r.DB.Create(&comment)
	return comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, commentID string) (string, error) {
	var comment *model.Comment

	r.DB.Where("id = ?", commentID).Delete(&comment)

	return "Comment verwijdert", nil
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

func (r *queryResolver) CommunityByID(ctx context.Context, communityID string) (*model.Community, error) {
	var community *model.Community

	r.DB.Where("id = ?", communityID).Find(&community)
	return community, nil
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
