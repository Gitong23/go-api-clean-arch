package controller

import (
	_inventoryService "github.com/Gitong23/go-api-clean-arch/pkg/inventory/service"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
}

func NewInventoryControllerImpl(inventoryService _inventoryService.InventoryService) InventroryController {
	return &inventoryControllerImpl{
		inventoryService,
	}
}
