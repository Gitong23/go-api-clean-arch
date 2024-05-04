package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type adminRepository interface {
	Creating(adminEntity *entities.Player) error
	FindByID(adminID string) (*entities.Player, error)
}
