package database

import (
	"context"
	"fmt"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url            = "mongodb://blackmarllbor0:funcConnect()@localhost:27017"
	Name           = "go-todo"
	CollectionName = "post"
)

var (
	ctx        = context.TODO()
	collection *mongo.Collection
	client     *mongo.Client
)

func ConnectDB() {
	// create client
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	// create connect
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to MongoDB")

	// check the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err, "error")
	}

	collection = client.Database(Name).Collection(CollectionName)
}

func Disconnect() {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}

func FindAll() []models.Post {
	var posts []models.Post                                       // создаем массив для заполнения
	cursor, err := collection.Find(ctx, bson.M{}, options.Find()) // получаем объект с данными
	if err != nil {
		log.Fatal(err)
	}
	// проходим по всем объектам
	for cursor.Next(ctx) {
		var post models.Post
		// декодируем данные
		if err := cursor.Decode(&post); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post) // заполняя данные

		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return posts
}

func InsertOne(post *models.Post) {
	if _, err := collection.InsertOne(ctx, post); err != nil {
		log.Fatal(err)
	}
}

func DeleteById(id string) {
	if _, err := collection.DeleteOne(ctx, bson.D{{"id", id}}, options.Delete()); err != nil {
		log.Fatal(err)
	}
}

func FindById(id string, post *models.Post) error {
	if err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&post); err != nil {
		return err
	}
	return nil
}

func UpdateById(id, title, content string) {
	option := bson.D{
		{"$set", bson.D{
			{"title", title},
			{"content", content},
		}},
	}
	if _, err := collection.UpdateOne(ctx, bson.D{{"id", id}}, option); err != nil {
		log.Fatal(err)
	}
}
