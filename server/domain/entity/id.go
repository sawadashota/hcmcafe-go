package entity

import "github.com/sawadashota/hcmcafe/server/lib/uuid"

type id struct {
	Id string `json:"id"`
}

func NewId(idStr string) *id {
	return &id{idStr}
}

func GenerateId() *id {
	return &id{uuid.Generate()}
}
