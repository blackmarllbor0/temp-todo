package main

import (
	"log"
	"os"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	h "github.com/blackmarllbor0/template_todo_server_in_go/handlers"

	"github.com/go-martini/martini"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/render"
)

func main() {
	// загрузка .env переменных
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	database.ConnectDB()        // подключение к базе данных
	defer database.Disconnect() // отключение от базы данных

	m := martini.Classic() // create new object
	// middleware для упрощения работы с html и json
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Get("/", h.IndexHandler) // главный обработчик

	m.Get("/write", h.WriterHandler)       // обработчик записи
	m.Post("/SavePost", h.SavePostHandler) // обработчик сохранения записей
	m.Get("/edit/:id", h.EditHandler)      // обработчик обновления данных в посте
	m.Get("/delete/:id", h.DeleteHandler)  // обработчик удаления поста

	m.Get("/login", h.GetLoginHandler)   // обработчик логина
	m.Post("/login", h.PostLoginHandler) // добовляем id сессии

	m.RunOnAddr(":" + os.Getenv("PORT")) // запуск сервера
}
