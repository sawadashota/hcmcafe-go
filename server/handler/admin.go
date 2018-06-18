package handler

import (
	"net/http"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
	"github.com/sawadashota/hcmcafe/server/domain/repository"
	"github.com/sawadashota/hcmcafe/server/domain/service"
)

type Admin struct{}

type FindAdminRequest struct {
	Id string `json:"id"`
}

type FindAdminResponse struct {
	Admin entity.Admin `json:"admin"`
}

func (a *Admin) Find(r *http.Request, args *FindAdminRequest, reply *FindAdminResponse) error {
	admin, err := repository.AdminRepository.Find(r, args.Id)

	if err != nil {
		return err
	}

	admin.HideCredentials()

	reply.Admin = *admin

	return nil
}

type CreateAdminRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
}

type CreateAdminResponse struct {
	Admin entity.Admin `json:"admin"`
}

func (a *Admin) Create(r *http.Request, args *CreateAdminRequest, reply *CreateAdminResponse) error {
	admin := entity.CreateAdmin(args.FirstName, args.LastName, args.Email, args.Password, args.Bio)

	err := repository.AdminRepository.Save(r, admin)

	if err != nil {
		return err
	}

	admin.HideCredentials()
	reply.Admin = *admin

	return nil
}

type AuthenticateAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateAdminResponse struct {
	Admin entity.Admin `json:"admin"`
}

func (a *Admin) Authenticate(r *http.Request, args *AuthenticateAdminRequest, reply *AuthenticateAdminResponse) error {
	admin, err := service.Authenticate(r, args.Email, args.Password)

	if err != nil {
		return err
	}

	reply.Admin = *admin

	return nil
}

type DeleteAdminRequest struct {
	Id string `json:"id"`
}

type DeleteAdminResponse struct{}

func (a *Admin) Delete(r *http.Request, args *DeleteAdminRequest, reply *DeleteAdminResponse) error {
	err := repository.AdminRepository.Delete(r, args.Id)

	if err != nil {
		return err
	}

	return nil
}

type GetAllAdminsRequest struct {
	Page int `json:"page"`
}

type GetAllAdminsResponse struct {
	Page        int            `json:"page"`
	EntirePage int            `json:"entire_page"`
	Admins      []entity.Admin `json:"admins"`
}

const PagePerData = 20

func (a *Admin) All(r *http.Request, args *GetAllAdminsRequest, reply *GetAllAdminsResponse) error {
	admins, entirePage, err := repository.AdminRepository.GetAll(r, PagePerData, args.Page)

	if err != nil {
		return err
	}

	reply.Page = args.Page
	reply.EntirePage = entirePage

	for _, admin := range admins {
		reply.Admins = append(reply.Admins, *admin)
	}

	return nil
}
