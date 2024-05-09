package controller

import (
	_playerCoinService "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/service"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{
		playerCoinService: playerCoinService,
	}
}
