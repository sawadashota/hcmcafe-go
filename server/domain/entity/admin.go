package entity

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Admin struct {
	Entity
	Name  Name   `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
	Bio   string `json:"bio"`
	password
	Token
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

func NewAdmin(id, firstName, lastName, email, encryptedPassword, bio, token string, createdAt, updatedAt, deletedAt time.Time) *Admin {
	return &Admin{
		Name:     *NewName(firstName, lastName),
		Email:    email,
		password: *NewPassword(encryptedPassword),
		Role:     *NewRole(),
		Bio:      bio,
		Token:    *NewToken(token),
		Entity:   *NewEntity(id, createdAt, updatedAt, deletedAt),
	}
}

func CreateAdmin(firstName, lastName, email, rawPassword, bio string) *Admin {
	p, _ := HashPassword(rawPassword)
	return &Admin{
		Name:     *NewName(firstName, lastName),
		Email:    email,
		password: *p,
		Role:     *NewRole(),
		Bio:      bio,
		Token:    *GenerateToken(),
		Entity:   *GenerateEntity(),
	}
}

func (a *Admin) Validate() error {
	if a.Id == "" {
		return fmt.Errorf("ID should be present\n")
	}

	if strings.Trim(a.Name.String(), " ") == "" {
		return fmt.Errorf("Name should be present\n")
	}

	if a.Email == "" {
		return fmt.Errorf("Email should be present\n")
	}

	if a.Password == "" {
		return fmt.Errorf("Password should be present\n")
	}

	if a.Role == "" {
		return fmt.Errorf("Role should be present\n")
	}

	return nil
}
