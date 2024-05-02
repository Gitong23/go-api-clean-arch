package controller

import (
	"net/http"

	"github.com/Gitong23/go-api-clean-arch/pkg/custom"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"

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
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList)
}
