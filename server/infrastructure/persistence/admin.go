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

//func (ar *AdminRepository) Find(ctx context.Context, id int) (*entity.Admin, error) {
//
//}

func (ar *AdminRepository) Save(r *http.Request, a *entity.Admin) error {
	return put(r, ar.kind, a)
}
