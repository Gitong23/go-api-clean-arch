package service

import (
	_playerCoinReposiotry "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/repository"
)

type playerCoinImpl struct {
	playerCoinRepository _playerCoinReposiotry.PlayerCoinRepository
}

func NewPlayerCoinService(playerCoinRepository _playerCoinReposiotry.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinImpl{
		playerCoinRepository: playerCoinRepository,
	}
}
