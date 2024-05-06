package service

import (
	_adminModel "github.com/Gitong23/go-api-clean-arch/pkg/admin/model"
	_playerModel "github.com/Gitong23/go-api-clean-arch/pkg/player/model"

	_adminRepository "github.com/Gitong23/go-api-clean-arch/pkg/admin/repository"
	_playerRepository "github.com/Gitong23/go-api-clean-arch/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	adminRepository _adminRepository.AdminRepository,
	playerRepository _playerRepository.PlayerRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAcountCreating(playerCreateReq *_playerModel.PlayerCreatingReq) error {
	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreateReq *_adminModel.AdminCreatingReq) error {
	return nil
}
