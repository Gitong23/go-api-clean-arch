package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
