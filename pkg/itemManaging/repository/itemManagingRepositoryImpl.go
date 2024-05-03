package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_itemManagingException "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemManagingRepositoryImp struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepository(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImp {
	return &itemManagingRepositoryImp{
		db:     db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImp) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Error creating item: %v", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImp) Editing(itemEntity *entities.Item) (*entities.Item, error) {

	// name := itemEntity.Name
	selectedID := itemEntity.ID
	item := new(entities.Item)
	if err := r.db.Model(item).Where("id = ?", selectedID).Updates(itemEntity).Error; err != nil {
		r.logger.Errorf("Error updating item: %v", err.Error())
		return nil, &_itemManagingException.ItemEditing{}
	}

	return item, nil
}
