package server

import (
	_inventoryRepository "github.com/Gitong23/go-api-clean-arch/pkg/inventory/repository"
	_itemShopController "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/controller"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"
	_playerCoinRepository "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {

	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepository(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		s.app.Logger,
	)

	itemShopController := _itemShopController.NewItemShopController(itemShopService)

	router.GET("", itemShopController.Listing, m.PlayerAuthorizing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)
}
