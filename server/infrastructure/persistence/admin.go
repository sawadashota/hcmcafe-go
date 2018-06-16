package persistence

import (
	"net/http"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
)

type AdminRepository datastoreRepository

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{
		kind: "Admin",
	}
}

// Find by ID(Key)
func (ar *AdminRepository) Find(r *http.Request, id string) (*entity.Admin, error) {
	a := new(entity.Admin)

	if err := find(r, ar.kind, id, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *AdminRepository) FindByEmail(r *http.Request, email string) (*entity.Admin, error) {
	a := new(entity.Admin)

	if err := first(r, ar.kind, "email", email, a); err != nil {
		return nil, err
	}
	return a, nil
}

// Save Admin Entity
func (ar *AdminRepository) Save(r *http.Request, a *entity.Admin) error {
	if err := a.Validate(); err != nil {
		return err
	}

	return put(r, ar.kind, a)
}
