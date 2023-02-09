package handlers

import (
	"net/http"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/blackmarllbor0/template_todo_server_in_go/models"
	"github.com/blackmarllbor0/template_todo_server_in_go/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// writerHandler парсит страницу ввода данных
func WriterHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, write, nil)
}

// savePostHandler создает и сохраняет полученные из ввода данные
func SavePostHandler(rnd render.Render, r *http.Request) {
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
func EditHandler(rnd render.Render, params martini.Params) {
	var post models.Post
	if err := database.FindById(params["id"], &post); err != nil {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(http.StatusOK, write, post) // передаем этот пост в date
}

// deleteHandler удаляет пост из хранилища
func DeleteHandler(rnd render.Render, params martini.Params) {
	if id := params["id"]; id == "" {
		rnd.Redirect("/")
		return
	} else {
		database.DeleteById(id)
		rnd.HTML(http.StatusOK, index, database.FindAll())
	}
}
