package controller

import (
	"net/http"

	"github.com/Gitong23/go-api-clean-arch/pkg/custom"
	_itemManagingModel "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/model"
	_itemMangingService "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type ItemManagingControllerImpl struct {
	itemManagingService _itemMangingService.ItemManagingService
}

func NewItemMannagingControllerImpl(
	itemManagingService _itemMangingService.ItemManagingService,
) ItemManagingController {
	return &ItemManagingControllerImpl{itemManagingService}
}

func (c *ItemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *ItemManagingControllerImpl) Editing(pctx echo.Context) error {
	itemEditingReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemEditingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Editting(itemEditingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, item)
}
