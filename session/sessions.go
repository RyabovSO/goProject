package sessions

import (
	"github.com/RyabovSO/goProject/utils"
	)

type sessionData struct {
	Username string
}

type Session struct {
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)	
	s.data = make (map[string]*sessionData)

	return s
}

func (s *Session) init(username string) string {
	sessionId := utils.GenerateId()
	data := &sessionData(Username: username)
	s.data[sessionId] = data

	return sessionId
}