package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
