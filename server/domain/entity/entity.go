package entity

import "time"

type entity struct {
	//id
	timestamps
}

func NewEntity(createdAt, updatedAt, deletedAt time.Time) *entity {
	return &entity{
		//id:         *NewId(id),
		timestamps: *NewTimestamp(createdAt, updatedAt, deletedAt),
	}
}

func GenerateEntity() *entity {
	return &entity{
		//id:         *GenerateId(),
		timestamps: *CreateTimestamp(),
	}
}
