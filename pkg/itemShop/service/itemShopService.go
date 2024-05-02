package service

import (
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
