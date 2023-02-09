package handlers

import "github.com/blackmarllbor0/template_todo_server_in_go/session"

const (
	index = "index" // home page
	write = "write" // write post page
	login = "login" // auth page

	COOKIE_NAME = "sessionId" // cookie name
)

var inMemorySsesion = session.NewSession() // храним сессию в оператианой паияти
