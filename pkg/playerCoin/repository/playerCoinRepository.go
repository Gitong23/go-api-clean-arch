package repository

import "github.com/Gitong23/go-api-clean-arch/entities"

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
}
