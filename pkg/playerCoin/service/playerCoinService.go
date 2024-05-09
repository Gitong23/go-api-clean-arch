package service

import (
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	Showing(playerID string) *_playerCoinModel.PlayerCoinShowing
}
