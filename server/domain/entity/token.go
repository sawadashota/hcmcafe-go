package entity

import (
	"time"

	"github.com/sawadashota/hcmcafe/server/lib/uuid"
)

type Session struct {
	Token     string    `json:"token,omitempty" datastore:"token"`
	UpdatedAt time.Time `json:"updated_at" datastore:"updated_at"`
}

func NewSession(tokenStr string, updatedAt time.Time) *Session {
	return &Session{
		Token:     tokenStr,
		UpdatedAt: updatedAt,
	}
}

func GenerateSession() *Session {
	return &Session{
		Token:     uuid.Generate(),
		UpdatedAt: time.Now(),
	}

}

func (s *Session) Refresh() {
	s.Token = uuid.Generate()
	s.UpdatedAt = time.Now()
}

func (s *Session) Flush() {
	s.Token = ""
	s.UpdatedAt = time.Now()
}
