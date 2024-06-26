package service

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
	_playerCoinModel "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/model"
	_playerCoinReposiotry "github.com/Gitong23/go-api-clean-arch/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _playerCoinReposiotry.PlayerCoinRepository
}

func NewPlayerCoinService(playerCoinRepository _playerCoinReposiotry.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{
		playerCoinRepository: playerCoinRepository,
	}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(nil, playerCoinEntity)
	if err != nil {
		return nil, err
	}
	playerCoinEntityResult.PlayerID = coinAddingReq.PlayerID

	return playerCoinEntityResult.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}
	return playerCoinShowing
}
