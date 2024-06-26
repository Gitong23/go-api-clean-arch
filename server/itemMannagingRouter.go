package server

import (
	_itemMannagingController "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/controller"
	_itemMannagingRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/repository"
	_itemMannagingService "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/service"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemMannagingRepository.NewItemManagingRepository(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemMannagingService.NewItemManagingService(itemManagingRepository, itemShopRepository)
	itemManagingController := _itemMannagingController.NewItemMannagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating, m.AdminAuthorizing)
	router.PATCH("/:itemID", itemManagingController.Editing, m.AdminAuthorizing)
	router.DELETE("/:itemID", itemManagingController.Archiving, m.AdminAuthorizing)
}
