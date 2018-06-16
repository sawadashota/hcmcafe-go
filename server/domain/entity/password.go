package entity

import "golang.org/x/crypto/bcrypt"

type password struct {
	Password string `json:"-" datastore:"password"`
}

func NewPassword(encryptedPassword string) *password {
	return &password{encryptedPassword}
}

func HashPassword(rawPassword string) (*password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &password{string(hash)}, nil
}

func (p *password) Verify(rawPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(rawPassword)) == nil
}
