package service

import (
	"net/http"
	"testing"

	"fmt"

	"time"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
	"github.com/sawadashota/hcmcafe/server/domain/repository"
	"github.com/sawadashota/hcmcafe/server/lib/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	repository.AdminRepository = &TestAdminRepository{}
	r := &http.Request{}

	cases := []struct {
		email    string
		password string
		expect   error
	}{
		{
			email:    "example@example.com",
			password: "1234",
			expect:   nil,
		},
		{
			email:    "hoge@example.com",
			password: "1234",
			expect:   fmt.Errorf("email: hoge@example.com is not found"),
		},
		{
			email:    "example@example.com",
			password: "12345",
			expect:   fmt.Errorf("failed to authenticate"),
		},
	}

	for _, c := range cases {
		_, err := Authenticate(r, c.email, c.password)

		if c.expect == nil {
			if err != nil {
				t.Errorf("%s", err)
			}
		} else {
			if err.Error() != c.expect.Error() {
				t.Errorf("%s", err)
			}
		}
	}
}

type TestAdminRepository struct{}

func (tar *TestAdminRepository) Save(r *http.Request, a *entity.Admin) error {
	return nil
}

func (tar *TestAdminRepository) Find(r *http.Request, id string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Id != id {
		return nil, fmt.Errorf("id: %s is not found", id)
	}

	return admin, nil
}

func (tar *TestAdminRepository) FindByEmail(r *http.Request, email string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Email != email {
		return nil, fmt.Errorf("email: %s is not found", email)
	}

	return admin, nil
}

func (tar *TestAdminRepository) FindByToken(r *http.Request, token string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Token.Token != token {
		return nil, fmt.Errorf("token: %s is not found", token)
	}

	return admin, nil
}

func adminMock() *entity.Admin {
	now := time.Time{}
	password, _ := bcrypt.Generate("1234")
	return entity.NewAdmin("123",
		"田中",
		"太郎",
		"example@example.com",
		password,
		"自己紹介させていただきます",
		"9876",
		now,
		now,
		now)
}
