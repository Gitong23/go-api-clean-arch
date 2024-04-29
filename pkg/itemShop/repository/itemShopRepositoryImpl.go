package repository

type itemShopRepositoryImpl struct{}

func NewItemShopRepository() ItemShopRepository {
	return &itemShopRepositoryImpl{}
}
