package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/blackmarllbor0/template_todo_server_in_go/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// template const's
const (
	index = "index"
	write = "write"
	login = "login"

	COOKIE_NAME = "sessionId"
)

// indexHandler парсит главную страницу
func indexHandler(rnd render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(r.Cookie(inMemorySsesion.Get(cookie.Value)))
		fmt.Println(cookie)
	}

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
		database.InsertOne(models.NewPost(utils.GenerateId(), title, content))
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

// getLoginHandler рендерит страницу входа
func getLoginHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, login, nil)
}

// postLoginHandler авторизует пользователя через сессию
func postLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	// password := r.FormValue("password")

	sessionId := inMemorySsesion.Init(username)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}
