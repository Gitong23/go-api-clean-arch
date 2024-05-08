package controller

import (
	"net/http"
	"strconv"

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
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *ItemManagingControllerImpl) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemEditingReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemEditingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Editing(itemID, itemEditingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *ItemManagingControllerImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}

func (c *ItemManagingControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}
	return itemIDUint64, nil
}
