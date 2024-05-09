package server

import (
	_playerCoinController "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/repository"
	_playerCoinService "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepository(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinService(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
}
