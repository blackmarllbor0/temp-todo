package main

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/db/documents"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

// template name const's
const (
	index = "index"
	write = "write"
)

// indexHandler парсит главную страницу
func indexHandler(rnd render.Render) {
	var postDoc []documents.Post
	if err := postsCollection.Find(nil).All(&postDoc); err != nil {
		log.Fatal(err)
	}

	var posts = make(models.Posts, len(postDoc))
	for _, doc := range postDoc {
		post := models.Post{Id: doc.Id, Title: doc.Title, Content: doc.Content}
		posts = append(posts, post)
	}

	rnd.HTML(http.StatusOK, index, posts)
}

// writerHandler парсит страницу ввода данных
func writerHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, write, nil)
}

// savePostHandler создает и сохраняет полученные из ввода данные
func savePostHandler(rnd render.Render, r *http.Request) {
	// получаем теги формы
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	postDoc := documents.Post{Id: id, Title: title, Content: content}
	// если id не имеет нулевое значение, то обновляем объект (/edit)
	if id != "" {
		if err := postsCollection.UpdateId(id, postDoc); err != nil {
			log.Fatal(err)
		}
	} else { // иначе создаем новый (/SavePost)
		postDoc.Id = GenerateId()
		if err := postsCollection.Insert(postDoc); err != nil {
			log.Fatal(err)
		}
	}

	// перенаправляем на страницу с уже созданными постами
	rnd.Redirect("/")
}

// editHandler обновляет запись в хранилище
func editHandler(rnd render.Render, params martini.Params) {
	id := params["id"] // получаем id элемента из request

	postDocument := documents.Post{}
	if err := postsCollection.FindId(id).One(&postDocument); err != nil {
		rnd.Redirect(index)
		return
	}

	post := models.Post{
		Id:      postDocument.Id,
		Title:   postDocument.Title,
		Content: postDocument.Content,
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

	// удаляем пост из базы
	if err := postsCollection.RemoveId(id); err != nil {
		log.Fatal(err)
	}

	rnd.Redirect("/")
}
