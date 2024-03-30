package types

import (
	"net/http"

	"github.com/google/uuid"
)

type Session struct {
	Session_id string
	Valid      bool
	User       User
}

var sessions = make(map[string]Session)

func NewSession(u *User) string {
	uuidStr := uuid.Must(uuid.NewRandom()).String()
	sessions[uuidStr] = Session{Session_id: uuidStr, Valid: true, User: *u}
	return uuidStr
}
func GetSession(r *http.Request) Session {
	cookie, err := r.Cookie("session")
	if err != nil {
		return Session{}
	}
	session, is := sessions[cookie.Value]
	if !is {
		return Session{}
	}
	return session
}

func NewSessionCookie(s string) *http.Cookie {
	return &http.Cookie{
		Name: "session", Value: s, HttpOnly: true, Secure: true, Path: "/", SameSite: http.SameSiteLaxMode,
	}
}
func (s *Session) RemoveSession() {
	delete(sessions, s.Session_id)
}
func (s *Session) Update() {
	sessions[s.Session_id] = *s
}
