// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
}

type Community struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	DescriptionHTML *string `json:"description_html"`
}

type NewCommunity struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
}

type Post struct {
	ID          string `json:"id"`
	CommunityID string `json:"community_id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentHTML string `json:"content_html"`
	Timestamp   int    `json:"timestamp"`
}

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Bio      *string `json:"bio"`
	Email    string  `json:"email"`
}

type UserCommunity struct {
	UserID      string `json:"user_id"`
	CommunityID string `json:"community_id"`
}