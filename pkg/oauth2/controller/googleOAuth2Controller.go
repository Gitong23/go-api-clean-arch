package controller

import (
	"github.com/Gitong23/go-api-clean-arch/config"
	_oauth2Service "github.com/Gitong23/go-api-clean-arch/pkg/oauth2/service"
	"github.com/labstack/echo/v4"
)

type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) OAuth2Controller {
	return &googleOAuth2Controller{
		oauth2Service,
		oauth2Conf,
		logger,
	}
}
