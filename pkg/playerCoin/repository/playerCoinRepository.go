package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
