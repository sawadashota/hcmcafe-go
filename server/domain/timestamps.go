package domain

import "time"

type timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// isNil judge timestamp is initialized
func isNil(timestamp *time.Time) bool {
	return timestamp.Unix() == time.Time{}.Unix()
}

func NewTimestamp(createdAt, updatedAt, deletedAt time.Time) *timestamps {
	return &timestamps{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func CreateTimestamp() *timestamps {
	return &timestamps{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}
}

func (t *timestamps) StampTime() {
	t.UpdatedAt = time.Now()
}

func (t *timestamps) Delete() {
	t.DeletedAt = time.Now()
}

func (t *timestamps) IsDeleted() bool {
	emptyTime := time.Time{}
	return t.DeletedAt != emptyTime
}
