package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_itemManagingModel "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
}
