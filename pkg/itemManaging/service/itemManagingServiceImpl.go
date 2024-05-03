package service

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_itemManagingModel "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/model"
	_itemManagingRepository "github.com/Gitong23/go-api-clean-arch/pkg/itemManaging/repository"
	_itemShopModel "github.com/Gitong23/go-api-clean-arch/pkg/itemShop/model"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingService(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {

	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editting(itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {

	itemEntity := &entities.Item{
		ID:          itemEditingReq.ID,
		Name:        itemEditingReq.Name,
		Description: itemEditingReq.Description,
		Picture:     itemEditingReq.Picture,
		Price:       itemEditingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Editing(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}
