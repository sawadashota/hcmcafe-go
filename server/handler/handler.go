package handler

import (
	"github.com/sawadashota/hcmcafe/server/domain/repository"
	"github.com/sawadashota/hcmcafe/server/infrastructure/persistence"
)

func init() {
	repository.AdminRepository = persistence.NewAdminRepository()
}
