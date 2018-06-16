package entity

import "time"

type Entity struct {
	Id
	timestamps
}

func NewEntity(id string, createdAt, updatedAt, deletedAt time.Time) *Entity {
	return &Entity{
		Id:         *NewId(id),
		timestamps: *NewTimestamp(createdAt, updatedAt, deletedAt),
	}
}

func GenerateEntity() *Entity {
	return &Entity{
		Id:         *GenerateId(),
		timestamps: *CreateTimestamp(),
	}
}
