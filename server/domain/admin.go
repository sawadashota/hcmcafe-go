package domain

import (
	"fmt"
	"net/url"
	"time"
)

type Admin struct {
	Name  Name   `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
	Bio   string `json:"bio"`
	password
	Token
	timestamps
}

type Avator struct {
	Url url.URL
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewName(firstName, lastName string) *Name {
	return &Name{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// String return full name
func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

func NewAdmin(firstName, lastName, email, encryptedPassword, bio, token string, createdAt, updatedAt, deletedAt time.Time) *Admin {
	return &Admin{
		Name:       *NewName(firstName, lastName),
		Email:      email,
		password:   *NewPassword(encryptedPassword),
		Role:       *NewRole(),
		Bio:        bio,
		Token:      *NewToken(token),
		timestamps: *NewTimestamp(createdAt, updatedAt, deletedAt),
	}
}

func CreateAdmin(firstName, lastName, email, rawPassword, bio string) *Admin {
	p, _ := HashPassword(rawPassword)
	return &Admin{
		Name:       *NewName(firstName, lastName),
		Email:      email,
		password:   *p,
		Role:       *NewRole(),
		Bio:        bio,
		Token:      *GenerateToken(),
		timestamps: *CreateTimestamp(),
	}
}
