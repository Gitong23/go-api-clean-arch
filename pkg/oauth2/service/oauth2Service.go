package service

import (
	_adminModel "github.com/Gitong23/go-api-clean-arch/pkg/admin/model"
	_playerModel "github.com/Gitong23/go-api-clean-arch/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAcountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}
