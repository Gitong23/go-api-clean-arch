package entities

import (
	"time"

	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
)

type PlayerCoin struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID  string    `gorm:"type:varchar(64);not null;"`
	Amount    int64     `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
}

func (p *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
	return &_playerCoinModel.PlayerCoin{
		ID:       int(p.ID),
		PlayerID: p.PlayerID,
		Amount:   int(p.Amount),
		CreateAt: p.CreatedAt,
	}
}
