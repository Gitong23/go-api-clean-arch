package server

import (
	_inventoryController "github.com/Gitong23/go-api-clean-arch/pkg/inventory/controller"
	_inventoryRepository "github.com/Gitong23/go-api-clean-arch/pkg/inventory/repository"
	_inventoryService "github.com/Gitong23/go-api-clean-arch/pkg/inventory/service"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(inventoryRepository, itemShopRepository)

	invertoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", invertoryController.Listing, m.PlayerAuthorizing)
}
