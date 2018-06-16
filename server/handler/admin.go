package handler

import (
	"net/http"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
	"github.com/sawadashota/hcmcafe/server/domain/repository"
)

type Admin struct{}

type FindAdminRequest struct {
	Id int `json:"id"`
}

type FindAdminResponse struct {
	Admin entity.Admin `json:"admin"`
	Error error        `json:"error"`
}

func (a *Admin) Find(r *http.Request, args *FindAdminRequest, reply *FindAdminResponse) error {

	return nil
}

type CreateAdminRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
}

type CreateAdminResponse struct{}

func (a *Admin) Create(r *http.Request, args *CreateAdminRequest, reply *CreateAdminResponse) error {
	admin := entity.CreateAdmin(args.FirstName, args.LastName, args.Email, args.Password, args.Bio)
	err := repository.AdminRepository.Save(r, admin)

	if err != nil {
		return err
	}

	return nil
}
