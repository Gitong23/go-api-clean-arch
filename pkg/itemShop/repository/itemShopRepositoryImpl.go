package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/execption"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepository(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {

	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %v", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
