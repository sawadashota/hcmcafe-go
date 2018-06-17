package entity

import "github.com/sawadashota/hcmcafe/server/lib/bcrypt"

type password struct {
	Password string `json:"-" datastore:"password"`
}

func NewPassword(encryptedPassword string) *password {
	return &password{encryptedPassword}
}

func HashPassword(rawPassword string) (*password, error) {
	hash, err := bcrypt.Generate(rawPassword)

	if err != nil {
		return nil, err
	}

	return &password{hash}, nil
}

func (p *password) Verify(rawPassword string) bool {
	return bcrypt.Compare(p.Password, rawPassword)
}
