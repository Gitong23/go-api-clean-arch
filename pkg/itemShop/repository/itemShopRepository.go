package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type ItemShopRepository interface{
	Listing() ([] *entities.Item, error)
}
