package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	"github.com/labstack/echo/v4"

	_itemShopException "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/execption"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepository(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {

	itemList := make([]*entities.Item, 0)

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false) // SELECT * FROM items

	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	// offset = (page -1) * size|limit
	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Errorf("Failed to list items: %v", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false) // SELECT * FROM items

	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting items failed: %v", err.Error())
		return -1, &_itemShopException.ItemListing{}
	}

	return count, nil
}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {

	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item: %v", err.Error())
		return nil, &_itemShopException.ItemNotFound{}
	}

	return item, nil
}

func (r *itemShopRepositoryImpl) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {

	items := make([]*entities.Item, 0)

	if err := r.db.Connect().Where("id in ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Errorf("Failed to find items: %v", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return items, nil
}
