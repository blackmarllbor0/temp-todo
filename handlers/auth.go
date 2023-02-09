package handlers

import (
	"net/http"
	"time"

	"github.com/martini-contrib/render"
)

// getLoginHandler рендерит страницу входа
func GetLoginHandler(rnd render.Render) {
	rnd.HTML(http.StatusOK, login, nil)
}

// postLoginHandler авторизует пользователя через сессию
func PostLoginHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	sessionId := inMemorySsesion.Init(username)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}
