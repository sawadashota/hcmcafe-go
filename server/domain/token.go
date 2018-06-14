package domain

import "github.com/google/uuid"

type Token struct {
	Token string `json:"Token"`
}

func NewToken(token string) *Token {
	return &Token{token}
}

func GenerateToken() *Token {
	return &Token{newUuid()}
}

func (t *Token) Refresh() {
	t.Token = newUuid()
}

func (t *Token) Flush() {
	t.Token = ""
}

func newUuid() string {
	return uuid.New().String()
}
