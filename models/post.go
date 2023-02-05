package models

type Post struct {
	Id      string
	Title   string
	Content string
}

type Posts []Post

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
