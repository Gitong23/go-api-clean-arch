package service

import (
	"github.com/Gitong23/go-api-clean-arch/entities"
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
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAcountCreating(playerCreateReq *_playerModel.PlayerCreatingReq) error {

	if !s.IsThisGuyIsReallyPlayer(playerCreateReq.ID) {
		playerEntity := &entities.Player{
			ID:     playerCreateReq.ID,
			Name:   playerCreateReq.Name,
			Email:  playerCreateReq.Email,
			Avatar: playerCreateReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreateReq *_adminModel.AdminCreatingReq) error {

	if !s.IsThisGuyIsReallyAdmin(adminCreateReq.ID) {
		adminEntity := &entities.Admin{
			ID:     adminCreateReq.ID,
			Name:   adminCreateReq.Name,
			Email:  adminCreateReq.Email,
			Avatar: adminCreateReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}

	}

	return nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}
	return player != nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}
	return admin != nil
}
