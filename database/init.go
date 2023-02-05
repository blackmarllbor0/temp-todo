package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Collection *mongo.Collection

func ConnectDB() *mongo.Database {
	connection := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("root")
}

//var Client *mongo.Client
//
//const url = "mongo://root:root@localhost:27017"
//
//// InitDB подключение к mongoDB
//func InitDB() {
//	// создаем контекст запуска
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	// подключаемся к БД
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	// пингуем, чтобы проверить подключение
//	if err := client.Ping(ctx, readpref.Primary()); err != nil {
//		log.Fatal(err)
//	}
//
//	// закрываем соединение с БД
//	defer func() {
//		if err = client.Disconnect(ctx); err != nil {
//			log.Fatal(err)
//		}
//	}()
//}
