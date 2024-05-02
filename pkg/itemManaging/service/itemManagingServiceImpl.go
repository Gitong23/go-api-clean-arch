package service

import (
	_itemManagingRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingService(
		itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	) ItemManagingService {
	return &itemManagingServiceImpl{ itemManagingRepository}
}

