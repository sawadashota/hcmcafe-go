package entity

import "github.com/sawadashota/hcmcafe/server/lib/uuid"

type Token struct {
	Token string `json:"Token"`
}

func NewToken(token string) *Token {
	return &Token{token}
}

func GenerateToken() *Token {
	return &Token{uuid.Generate()}
}

func (t *Token) Refresh() {
	t.Token = uuid.Generate()
}

func (t *Token) Flush() {
	t.Token = ""
}
