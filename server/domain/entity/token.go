package entity

import "github.com/sawadashota/hcmcafe/server/lib/uuid"

type Token struct {
	Token string `json:"token,omitempty" datastore:"token"`
}

func NewToken(tokenStr string) *Token {
	return &Token{tokenStr}
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
