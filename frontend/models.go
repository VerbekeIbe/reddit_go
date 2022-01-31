package main

type Community struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type PostObject struct {
	Id string
	Title string
	Content string
	Timestamp int64
	CommunityId string `json:"community_id"`
}

type CommentObject struct {
	Id string
	Content string
	Timestamp int64
	UserId string `json:"user_id"`
}

type NewUserCommunity struct {
	UserId string `json:"user_id"`
	CommunityId string `json:"community_id"`
}

type NewPost struct {
	UserId string `json:"user_id"`
	CommunityId string `json:"community_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ContentHTML string `json:"content_html"`
}

type NewComment struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
	Content string `json:"content"`
}
