package entity

import "time"

type Entity struct {
	id
	timestamps
}

func NewEntity(id string, createdAt, updatedAt, deletedAt time.Time) *Entity {
	return &Entity{
		id:         *NewId(id),
		timestamps: *NewTimestamp(createdAt, updatedAt, deletedAt),
	}
}

func GenerateEntity() *Entity {
	return &Entity{
		id:         *GenerateId(),
		timestamps: *CreateTimestamp(),
	}
}
