package session

import (
	"github.com/blackmarllbor0/template_todo_server_in_go/utils"
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

const CookieName = "sessionId"

type Session struct {
	id       string
	Username string
	IsAuth   bool
}

type SessionStore struct {
	data map[string]*Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		data: make(map[string]*Session),
	}
}

func (s *SessionStore) Get(sessionId string) *Session {
	session := s.data[sessionId]
	if session != nil {
		return &Session{id: sessionId}
	}
	return session
}

func (s *SessionStore) Set(session *Session) {
	s.data[session.id] = session
}

func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(CookieName)
	if cookie != nil {
		return cookie.Value
	}
	sessionId := utils.GenerateId()

	cookie = &http.Cookie{
		Name:    CookieName,
		Value:   sessionId,
		Expires: time.Now().Add(time.Minute * 5),
	}

	return sessionId
}

var sessionStore = NewSessionStore()

func Middleware(ctx martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := ensureCookie(r, w)
	session := sessionStore.Get(sessionId)

	ctx.Map(sessionId)
	ctx.Next()

	sessionStore.Set(session)
}
