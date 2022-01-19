// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// DATABASE TYPES

type User struct {
	ID       string  `json:"id" gorm:"primary_key;size:256"`
	Username string  `json:"username"`
	Bio      *string `json:"bio"`
	Email    string  `json:"email"`
}

type Community struct {
	ID              string  `json:"id" gorm:"primary_key;size:256"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	DescriptionHTML *string `json:"description_html"`
}

type Post struct {
	ID          string `json:"id" gorm:"primary_key;size:256"`
	CommunityID string `json:"community_id"`
	Community Community
	UserID      string `json:"user_id"`
	User User
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentHTML string `json:"content_html"`
	Timestamp   int    `json:"timestamp"`
}

type Comment struct {
	ID        string `json:"id" gorm:"primary_key;size:256"`
	PostID    string `json:"post_id"`
	Post Post
	UserID    string `json:"user_id"`
	User User
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
}

type UserCommunity struct {
	UserID      string `json:"user_id"`
	User User
	CommunityID string `json:"community_id"`
	Community Community
}

// MUTATION TYPES

type NewCommunity struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
}
