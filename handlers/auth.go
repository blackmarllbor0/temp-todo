package handlers

import (
	"net/http"

	"github.com/blackmarllbor0/template_todo_server_in_go/session"
	"github.com/martini-contrib/render"
)

// getLoginHandler рендерит страницу входа
func GetLoginHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, login, nil)
}

// postLoginHandler авторизует пользователя через сессию
func PostLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request, session *session.Session) {
	username := r.FormValue("username")

	session.Username = username
	session.IsAuth = true

	rnd.Redirect("/")
}
