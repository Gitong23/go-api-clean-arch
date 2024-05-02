package controller

import (
	_itemMangingService "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/service"
)

type ItemManagingControllerImpl struct {
	itemManagingService _itemMangingService.ItemManagingService
}

func NewItemMannagingControllerImpl(
	itemManagingService _itemMangingService.ItemManagingService,
) ItemManagingController {
	return &ItemManagingControllerImpl{itemManagingService}
}
