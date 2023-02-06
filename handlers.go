package main

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

// template const's
const (
	index = "index"
	write = "write"
)

// indexHandler парсит главную страницу
func indexHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, index, database.FindAll())
}

// writerHandler парсит страницу ввода данных
func writerHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, write, nil)
}

// savePostHandler создает и сохраняет полученные из ввода данные
func savePostHandler(rnd render.Render, r *http.Request) {
	// получаем теги формы
	title := r.FormValue("title")
	content := r.FormValue("content")

	if id := r.FormValue("id"); id != "" {
		// если id не имеет нулевое значение, то обновляем объект (/edit)
		database.UpdateById(id, title, content)
	} else {
		// иначе создаем новый (/SavePost)
		database.InsertOne(models.NewPost(GenerateId(), title, content))
	}

	// перенаправляем на страницу с уже созданными постами
	rnd.Redirect("/")
}

// editHandler обновляет запись в хранилище
func editHandler(rnd render.Render, params martini.Params) {
	var post models.Post
	if err := database.FindById(params["id"], &post); err != nil {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(http.StatusOK, write, post) // передаем этот пост в date
}

// deleteHandler удаляет пост из хранилища
func deleteHandler(rnd render.Render, params martini.Params) {
	if id := params["id"]; id == "" {
		rnd.Redirect("/")
		return
	} else {
		database.DeleteById(id)
		rnd.HTML(http.StatusOK, index, database.FindAll())
	}
}
