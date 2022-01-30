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
