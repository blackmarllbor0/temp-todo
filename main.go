package main

import (
	"fmt"
	"github.com/blackmarllbor0/todo_go/models"
	"log"
	"net/http"
)

// server const's
const (
	serverHost = "http://localhost"
	serverPort = ":8080"
)

// Временное хранилище данных
var posts models.Posts

func main() {
	// вывод
	fmt.Println("The server is listen on", serverHost+serverPort)

	// инициализируем хранилище
	posts = make(map[string]*models.Post, 0)

	// обработка создания стилей в index.html из папки assets
	//http.Handle(assets, http.StripPrefix(assets, http.FileServer(http.Dir("."+assets))))
	
	// обработчики путей
	http.HandleFunc("/", indexHandler)            // главный обработчик
	http.HandleFunc("/write", writerHandler)      // обработчик записи
	http.HandleFunc("/SavePost", savePostHandler) // обработчик сохранения записей
	http.HandleFunc("/edit", editHandler)         // обработчик обновления данных в посте
	http.HandleFunc("/delete", deleteHandler)     // обработчик удаления поста

	// запуск сервера
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal("There was an error when starting the server\n", err)
	}
}
