package main

import (
	"context"
	"fmt"
	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// server const's
const serverPort = ":8080"

// Временное хранилище данных
var posts models.Posts

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func main() {
	client := database.ConnectDB(ctx)
	collection = client.Database(database.DatabaseName).Collection(database.CollectionName)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	insertResult, err := collection.InsertOne(ctx, models.Post{
		Id:      "1",
		Title:   "2",
		Content: "3",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult)

	// инициализируем хранилище
	posts = make(map[string]*models.Post, 0)

	m := martini.Classic() // create new object
	// middleware для упрощения работы с html и json
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		//Funcs:      []template.FuncMap(nil),
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	//m.Use(martini.Static("assets", martini.StaticOptions{Prefix: "assets"}))

	// обработчики путей
	m.Get("/", indexHandler)             // главный обработчик
	m.Get("/write", writerHandler)       // обработчик записи
	m.Post("/SavePost", savePostHandler) // обработчик сохранения записей
	m.Get("/edit/:id", editHandler)      // обработчик обновления данных в посте
	m.Get("/delete/:id", deleteHandler)  // обработчик удаления поста

	// запуск сервера
	m.RunOnAddr(serverPort)
}
