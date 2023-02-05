package documents

type Post struct {
	Id      string `bson:"_id,omitempty"`
	Title   string
	Content string
}
