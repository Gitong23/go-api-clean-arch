package controller

import (
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopController(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
