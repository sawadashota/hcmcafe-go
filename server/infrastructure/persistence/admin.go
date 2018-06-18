package persistence

import (
	"net/http"

	"fmt"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
)

type AdminRepository datastoreRepository

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{
		kind: "Admin",
	}
}

// Find by ID(Key)
func (ar *AdminRepository) GetAll(r *http.Request, limit, page int) ([]*entity.Admin, int, error) {
	var as []*entity.Admin
	entirePage, err := getAll(r, ar.kind, limit, page)(&as)

	return as, entirePage, err
}

// Find by ID(Key)
func (ar *AdminRepository) Find(r *http.Request, id string) (*entity.Admin, error) {
	a := new(entity.Admin)
	a.Id = entity.NewId(id)

	if err := find(r, id, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *AdminRepository) FindByEmail(r *http.Request, email string) (*entity.Admin, error) {
	a := new(entity.Admin)

	if err := first(r, "email", email, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *AdminRepository) FindByToken(r *http.Request, token string) (*entity.Admin, error) {
	a := new(entity.Admin)

	if err := first(r, "token", token, a); err != nil {
		return nil, err
	}
	return a, nil
}

// Save Admin Entity
func (ar *AdminRepository) Save(r *http.Request, a *entity.Admin) error {
	if err := a.Validate(); err != nil {
		return err
	}

	if err := ar.uniqueCheck(r, a); err != nil {
		return err
	}

	return put(r, a)
}

func (ar *AdminRepository) Delete(r *http.Request, id string) error {
	a := new(entity.Admin)
	a.Id = entity.NewId(id)

	if err := destroy(r, a); err != nil {
		return err
	}

	return nil
}

func (ar *AdminRepository) uniqueCheck(r *http.Request, a *entity.Admin) error {
	items := []struct {
		key   string
		value interface{}
	}{
		{
			key:   "email",
			value: a.Email,
		},
	}

	for _, i := range items {
		if e, err := exist(r, ar.kind, a.Id.String(), i.key, i.value); err != nil {
			return err
		} else if e {
			return fmt.Errorf("%v is already exist", i.value)
		}
	}

	return nil
}
