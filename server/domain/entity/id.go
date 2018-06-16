package entity

import "github.com/sawadashota/hcmcafe/server/lib/uuid"

type Id struct {
	Id string `json:"Id"`
}

func NewId(id string) *Id {
	return &Id{id}
}

func GenerateId() *Id {
	return &Id{uuid.Generate()}
}
