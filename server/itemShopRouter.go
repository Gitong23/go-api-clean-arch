package server

import (
	_itemShopController "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/controller"
	_itemShopRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/repository"
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {

	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepository(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopController(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
