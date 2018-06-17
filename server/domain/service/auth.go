package service

import (
	"net/http"

	"fmt"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
	"github.com/sawadashota/hcmcafe/server/domain/repository"
)

func Authenticate(r *http.Request, email, password string) (*entity.Admin, error) {
	admin, err := repository.AdminRepository.FindByEmail(r, email)

	if err != nil {
		return nil, err
	}

	if !admin.Verify(password) {
		return nil, fmt.Errorf("failed to authenticate")
	}

	admin.Session.Refresh()

	repository.AdminRepository.Save(r, admin)

	return admin, nil
}

func AuthUser(r *http.Request, token string) (*entity.Admin, error) {
	admin, err := repository.AdminRepository.FindByToken(r, token)

	if err != nil {
		return nil, err
	}

	return admin, nil
}
