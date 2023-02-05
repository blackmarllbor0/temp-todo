package models

type Post struct {
	Id      string `bson:"_id"`
	Title   string
	Content string
}

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}

type Posts map[string]*Post
