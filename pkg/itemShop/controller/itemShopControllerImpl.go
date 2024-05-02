package controller

import (
	"net/http"

	"github.com/Gitong23/go-api-clean-arch/pkg/custom"
	_itemShopService "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopController(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		// return pctx.String(http.StatusInternalServerError, err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList)
}
