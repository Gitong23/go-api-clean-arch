package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	_itemManagingException "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/exception"
	_itemManagingModel "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/model"
	"github.com/labstack/echo/v4"
)

type itemManagingRepositoryImp struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepository(db databases.Database, logger echo.Logger) *itemManagingRepositoryImp {
	return &itemManagingRepositoryImp{
		db:     db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImp) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Error creating item: %v", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImp) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImp) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
