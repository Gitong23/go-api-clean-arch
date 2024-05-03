package service

import (
	_itemManagingModel "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/model"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
	Archiving(itemID uint64) error
}
