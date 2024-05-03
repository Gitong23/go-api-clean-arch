package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
}
