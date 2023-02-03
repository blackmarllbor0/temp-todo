package main

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"net/http"
)

// template const's
const (
	index  = "templates/index.html"
	header = "templates/header.html"
	footer = "templates/footer.html"
	writer = "templates/write.html"
	assets = "/assets/"
)

// indexHandler парсит главную страницу
func indexHandler(w http.ResponseWriter, _ *http.Request) {
	generateTemplate(w, index, "index", posts)
}

// writerHandler парсит страницу ввода данных
func writerHandler(w http.ResponseWriter, _ *http.Request) {
	generateTemplate(w, writer, "write", nil)
}

// savePostHandler создает и сохраняет полученные из ввода данные
func savePostHandler(w http.ResponseWriter, r *http.Request) {
	// получаем теги формы
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post

	// если id не имеет нулевое значение, то обновляем объект (/edit)
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else { // иначе создаем новый (/SavePost)
		id = GenerateId()
		post = models.NewPost(id, title, content)
		posts[post.Id] = post
	}

	// перенаправляем на страницу с уже созданными постами
	http.Redirect(w, r, "/", http.StatusFound)
}

// editHandler обновляет запись в хранилище
func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("Id")  // получаем id элемента из request
	post, found := posts[id] // ищем в базе по id
	if !found {
		http.NotFound(w, r)
	}

	generateTemplate(w, writer, "write", post) // передаем этот пост в date
}

// deleteHandler удаляет пост из хранилища
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("Id")
	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id) // удаляем пост из базы

	generateTemplate(w, index, "index", posts)
}
