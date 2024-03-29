package repository

import (
	"context"

	"net/http"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
)

var (
	AdminRepository adminRepository
	//CafeRepository  cafeRepository
)

type adminRepository interface {
	Save(r *http.Request, a *entity.Admin) error
	Delete(r *http.Request, id string) error
	GetAll(r *http.Request, limit, page int) ([]*entity.Admin, int, error)
	Find(r *http.Request, id string) (*entity.Admin, error)
	FindByEmail(r *http.Request, email string) (*entity.Admin, error)
	FindByToken(r *http.Request, token string) (*entity.Admin, error)
}

type cafeRepository interface {
	Save(ctx context.Context, cafe *entity.Cafe) error
	Find(ctx context.Context, id int) (*entity.Cafe, error)
	Get() ([]*entity.Cafe, error)
}
