package handlers

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/database/models"
	"github.com/blackmarllbor0/template_todo_server_in_go/session"
	"net/http"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/martini-contrib/render"
)

// IndexHandler парсит главную страницу
func IndexHandler(rnd render.Render, r *http.Request, session *session.Session) {
	model := models.PostList{}
	model.IsAuth = session.IsAuth
	model.Posts = database.FindAll()

	rnd.HTML(http.StatusOK, index, database.FindAll())
}
