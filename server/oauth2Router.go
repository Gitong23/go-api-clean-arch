package server

import (
	_adminRepository "github.com/Gitong23/go-api-clean-arch/pkg/admin/repository"
	_oauth2Controller "github.com/Gitong23/go-api-clean-arch/pkg/oauth2/controller"
	_oauth2Service "github.com/Gitong23/go-api-clean-arch/pkg/oauth2/service"
	_playerRepository "github.com/Gitong23/go-api-clean-arch/pkg/player/repository"
)

func (s *echoServer) initOAuth2Router() {
	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(oauth2Service, s.conf.OAuth2, s.app.Logger)

	// router.GET("/test", oauth2Controller.Test)
	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/player/login/callback", oauth2Controller.PlayerLoginCallback)
	router.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	router.DELETE("/logout", oauth2Controller.Logout)
}
