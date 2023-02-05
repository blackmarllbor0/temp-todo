package main

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

// template const's
const (
	index = "index"
	write = "write"
)

// indexHandler парсит главную страницу
func indexHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, index, posts)
}

// writerHandler парсит страницу ввода данных
func writerHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, write, nil)
}

// savePostHandler создает и сохраняет полученные из ввода данные
func savePostHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
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
		if _, err := collection.InsertOne(ctx, post); err != nil {
			log.Fatal(err)
		}
	}

	// перенаправляем на страницу с уже созданными постами
	rnd.Redirect("/")
}

// editHandler обновляет запись в хранилище
func editHandler(rnd render.Render, params martini.Params) {
	id := params["id"] // получаем id элемента из request

	post, found := posts[id] // ищем в базе по id
	if !found {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(http.StatusOK, write, post) // передаем этот пост в date
}

// deleteHandler удаляет пост из хранилища
func deleteHandler(rnd render.Render, params martini.Params) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	delete(posts, id) // удаляем пост из базы

	rnd.HTML(http.StatusOK, index, posts)
}
