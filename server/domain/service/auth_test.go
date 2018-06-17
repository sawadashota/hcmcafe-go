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
	repository.AdminRepository = &testAdminRepository{}
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

type testAdminRepository struct{}

func (tar *testAdminRepository) Save(r *http.Request, a *entity.Admin) error {
	return nil
}

func (tar *testAdminRepository) Find(r *http.Request, id string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Id != id {
		return nil, fmt.Errorf("id: %s is not found", id)
	}

	return admin, nil
}

func (tar *testAdminRepository) FindByEmail(r *http.Request, email string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Email != email {
		return nil, fmt.Errorf("email: %s is not found", email)
	}

	return admin, nil
}

func (tar *testAdminRepository) FindByToken(r *http.Request, token string) (*entity.Admin, error) {
	admin := adminMock()

	if admin.Session.Token != token {
		return nil, fmt.Errorf("token: %s is not found", token)
	}

	return admin, nil
}

func (tar *testAdminRepository) Delete(r *http.Request, id string) error {
	admin := adminMock()

	if admin.Id != id {
		return fmt.Errorf("id: %s is not found", id)
	}

	return nil
}

func adminMock() *entity.Admin {
	now := time.Time{}
	password, _ := bcrypt.Generate("1234")
	a := entity.NewAdmin("123",
		"田中",
		"太郎",
		"example@example.com",
		password,
		"自己紹介させていただきます",
		now,
		now,
		now)

	a.Session = *entity.NewSession("9876", time.Now())

	return a
}
