package handlers

import (
	"fmt"
	"net/http"

	"github.com/blackmarllbor0/template_todo_server_in_go/database"
	"github.com/martini-contrib/render"
)

// indexHandler парсит главную страницу
func IndexHandler(rnd render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(r.Cookie(inMemorySsesion.Get(cookie.Value)))
		fmt.Println(cookie)
	}

	rnd.HTML(http.StatusOK, index, database.FindAll())
}
