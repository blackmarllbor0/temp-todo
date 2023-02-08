package session

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/utils"
)

type Data struct {
	Username string
}

type Session struct {
	data map[string]*Data
}

func NewSession() *Session {
	s := new(Session)
	s.data = make(map[string]*Data)

	return s
}

func (s *Session) init(username string) string {
	sessionId := utils.GenerateId()

	data := &Data{username}
	s.data[sessionId] = data

	return sessionId
}
