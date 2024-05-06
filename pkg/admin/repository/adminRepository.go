package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}
