package server

import (
	_itemMannagingController "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/controller"
	_itemMannagingRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/repository"
	_itemMannagingService "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemMannagingRepository.NewItemManagingRepository(s.db, s.app.Logger)
	itemManagingService := _itemMannagingService.NewItemManagingService(itemManagingRepository)
	itemManagingController := _itemMannagingController.NewItemMannagingControllerImpl(itemManagingService)

	_ = itemManagingController
	_ = router
}
