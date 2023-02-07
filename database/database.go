package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	url          = "mongodb://localhost:27017"
	DatabaseName = "go-todo"

	CollectionName = "post"
)

func ConnectDB(ctx context.Context) *mongo.Client {
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

	return client
}
