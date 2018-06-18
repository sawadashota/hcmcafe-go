package entity

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Admin struct {
	Id Id `json:"id" datastore:"-" goon:"id"`
	entity
	Name  Name   `json:"name" datastore:"name,noindex"`
	Email string `json:"email" datastore:"email"`
	Role  Role   `json:"role" datastore:"role"`
	Bio   string `json:"bio" datastore:"bio,noindex"`
	password
	Session Session `json:"session,omitempty" datastore:"session"`
}

type Avator struct {
	Url url.URL
}

type Name struct {
	FirstName string `json:"first_name" datastore:"first_name"`
	LastName  string `json:"last_name" datastore:"last_name"`
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

func NewAdmin(id, firstName, lastName, email, encryptedPassword, bio string, createdAt, updatedAt, deletedAt time.Time) *Admin {
	return &Admin{
		Id:       NewId(id),
		Name:     *NewName(firstName, lastName),
		Email:    email,
		password: *NewPassword(encryptedPassword),
		Role:     *NewRole(),
		Bio:      bio,
		entity:   *NewEntity(createdAt, updatedAt, deletedAt),
	}
}

func CreateAdmin(firstName, lastName, email, rawPassword, bio string) *Admin {
	p, _ := HashPassword(rawPassword)
	return &Admin{
		Id:       GenerateId(),
		Name:     *NewName(firstName, lastName),
		Email:    email,
		password: *p,
		Role:     *NewRole(),
		Bio:      bio,
		Session:  *GenerateSession(),
		entity:   *GenerateEntity(),
	}
}

func (a *Admin) Validate() error {
	if a.Id == "" {
		return fmt.Errorf("ID should be present")
	}

	if strings.Trim(a.Name.String(), " ") == "" {
		return fmt.Errorf("name should be present")
	}

	if a.Email == "" {
		return fmt.Errorf("email should be present")
	}

	if a.Password == "" {
		return fmt.Errorf("password should be present")
	}

	if a.Role == "" {
		return fmt.Errorf("role should be present")
	}

	return nil
}

// HideCredentials hide important data
func (a *Admin) HideCredentials() {
	a.Session = Session{}
}
