package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type InventoryRepository interface {
	Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
