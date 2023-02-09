package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blackmarllbor0/template_todo_server_in_go/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx        = context.TODO()
	collection *mongo.Collection
	client     *mongo.Client
)

// ConnectDB подключается к базе данных
func ConnectDB() {
	// create client
	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectString()))
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

	// подключаемся к конкретной таблице и коллекции
	collection = client.Database("go-todo").Collection("post")
}

// getConnectString возвращает строку подключения MongoDB
func getConnectString() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@localhost:27017",
		os.Getenv("MONGO_USER_NAME"),
		os.Getenv("MONGO_USER_PASSWORD"),
	)
}

// Disconnect отключается от базы данных
func Disconnect() {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}

// FindAll ищет все посты в коллекции и возвращает уже преобразованные данные
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

// InsertOne записывает новый пост в коллекцию
func InsertOne(post *models.Post) {
	if _, err := collection.InsertOne(ctx, post); err != nil {
		log.Fatal(err)
	}
}

// DeleteById удаляет пост по id
func DeleteById(id string) {
	if _, err := collection.DeleteOne(ctx, bson.D{{"id", id}}, options.Delete()); err != nil {
		log.Fatal(err)
	}
}

// FindById ищет пост по id, и получает ссылку на post, и преобразует этот пост в нужный вид
func FindById(id string, post *models.Post) error {
	if err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&post); err != nil {
		return err
	}
	return nil
}

// UpdateById обновляет пост по id. Принимает только title и content
func UpdateById(id, title, content string) {
	option := bson.D{ // настройки для изменения
		{"$set", bson.D{
			{"title", title},
			{"content", content},
		}},
	}
	if _, err := collection.UpdateOne(ctx, bson.D{{"id", id}}, option); err != nil {
		log.Fatal(err)
	}
}
