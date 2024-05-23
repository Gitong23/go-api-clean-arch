package controller

import (
	"fmt"
	"net/http"

	"github.com/Gitong23/go-api-clean-arch/pkg/custom"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"
	"github.com/Gitong23/go-api-clean-arch/pkg/validation"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopController(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {

	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	fmt.Printf("$$$$$$$$$itemModelList: %s\n", itemModelList.Items[0].Name)
	return pctx.JSON(http.StatusOK, itemModelList)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	buyingReq := new(_itemShopModel.BuyingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, playerCoin)

}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	return nil
}
