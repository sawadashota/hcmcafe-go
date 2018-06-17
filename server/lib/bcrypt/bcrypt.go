package bcrypt

import "golang.org/x/crypto/bcrypt"

func Generate(rawPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(encryptedPassword, rawPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(rawPassword)) == nil
}
