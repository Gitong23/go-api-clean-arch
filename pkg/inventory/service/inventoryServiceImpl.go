package service

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_inventoryModel "github.com/Gitong23/go-api-clean-arch/pkg/inventory/model"
	_inventoryRepository "github.com/Gitong23/go-api-clean-arch/pkg/inventory/repository"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(
	inventoryRepository _inventoryRepository.InventoryRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		itemShopRepository:  itemShopRepository,
	}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {

	inventoryEntities, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqeItemWithQuantityCounterList(inventoryEntities)

	return s.buildInventoryListingResult(
		uniqueItemWithQuantityCounterList,
	), nil
}

func (s *inventoryServiceImpl) buildInventoryListingResult(
	uqiqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []*_inventoryModel.Inventory {

	uniqueItemIDList := s.getItemID(uqiqueItemWithQuantityCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uqiqueItemWithQuantityCounterList)

	for _, itemEntitie := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item:     itemEntitie.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntitie.ID],
		})
	}

	return results
}

func (s *inventoryServiceImpl) getItemID(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getUniqeItemWithQuantityCounterList(
	inventoryEntities []*entities.Inventory,
) []_inventoryModel.ItemQuantityCounting {

	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
