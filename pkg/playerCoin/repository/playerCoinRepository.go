package repository

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
