package models

type Post struct {
	Id      string
	Title   string
	Content string
}

type Posts map[string]*Post

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
