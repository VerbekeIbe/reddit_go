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
}