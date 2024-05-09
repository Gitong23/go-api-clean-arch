package service

import (
	_invertoryModel "github.com/Gitong23/go-api-clean-arch/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_invertoryModel.Inventory, error)
}
