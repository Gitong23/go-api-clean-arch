package service

import (
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
}
