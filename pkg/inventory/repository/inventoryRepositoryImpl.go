package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	_inventoryException "github.com/Gitong23/go-api-clean-arch/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {

	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(
		inventoryEntities,
	).Error; err != nil {
		r.logger.Errorf("error filling inventory: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {

	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities, err := r.findPlayerItemInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := conn.Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s", err.Error())
			return &_inventoryException.PlayerItemRemoving{
				ItemID: itemID,
			}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {

	var count int64

	if err := r.db.Connect().Model(
		&entities.Inventory{},
	).Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Count(
		&count,
	).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory: %s", err.Error())
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {

	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and is_deleted = ?", playerID, false,
	).Find(
		&inventoryEntities,
	).Error; err != nil {
		r.logger.Errorf("error listing player item in inventory: %s", err.Error())
		return nil, &_inventoryException.PlayerItemsFinging{
			PlayerID: playerID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) findPlayerItemInventoryByID(
	playerID string,
	itemID uint64,
	limit int,
) ([]*entities.Inventory, error) {

	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error finding player item in inventory by ID: %s", err.Error())
		return nil, &_inventoryException.PlayerItemRemoving{
			ItemID: itemID,
		}
	}

	return inventoryEntities, nil
}
