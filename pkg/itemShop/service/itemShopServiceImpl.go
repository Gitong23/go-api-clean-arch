package service

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_inventoryRepository "github.com/Gitong23/go-api-clean-arch/pkg/inventory/repository"
	_itemShopException "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/execption"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/repository"
	"github.com/labstack/echo/v4"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	logger echo.Logger,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		logger,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCouting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCouting, int64(itemFilter.Size))
	return s.toItemResultResponse(itemList, int64(itemFilter.Page), totalPage), nil
}

// 1. Find Item by ID
// 2. Total Price Calculation
// 3. Check Player Coin
// 4. Purchase History Recording
// 5. Coin Deducting
// 6. Inventory Filling
// 7. Return Player Coin
func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), buyingReq.Quantity)

	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
	})
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Infof("Purchase History Recording Success: %s", purchaseRecording.ID)

	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})

	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}

	s.logger.Infof("Player coin deducted: %d", playerCoin.Amount)

	inventoryEntity, err := s.inventoryRepository.Filling(
		tx,
		buyingReq.PlayerID,
		buyingReq.ItemID,
		int(buyingReq.Quantity),
	)
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		return nil, err
	}
	s.logger.Infof("Inventory filled: %d", len(inventoryEntity))

	return playerCoin.ToPlayerCoinModel(), nil
}

// 1. Find Item by ID
// 2. Total Price Calculation
// 3. Check Player Item
// 4. Purchase History Recording
// 5. Coin Adding
// 6. Inventory Removing
// 7. Return Player Coin
func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	return nil, nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size > 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Error("Player Coin Not Enough")
		return &_itemShopException.CoinNotEnough{}
	}
	return nil
}
