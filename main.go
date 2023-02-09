package main

import (
	"log"
	"os"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/blackmarllbor0/template_todo_server_in_go/session"

	"github.com/go-martini/martini"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/render"
)

var inMemorySsesion *session.Session // храним сессию в оператианой паияти

func main() {
	// загрузка .env переменных
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	inMemorySsesion = session.NewSession() // create a new session

	// подключение к базе данных
	database.ConnectDB()
	// отключение от базы данных
	defer database.Disconnect()

	m := martini.Classic() // create new object
	// middleware для упрощения работы с html и json
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
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
	m.Get("/login", getLoginHandler)     // обработчик логина
	m.Post("/login", postLoginHandler)   // добовляем id сессии

	// запуск сервера
	m.RunOnAddr(":" + os.Getenv("PORT"))
}
