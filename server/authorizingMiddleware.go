package server

import (
	"github.com/Gitong23/go-api-clean-arch/config"
	_oauth2Controller "github.com/Gitong23/go-api-clean-arch/pkg/oauth2/controller"
	"github.com/labstack/echo/v4"
)

type authorizingMiddleware struct {
	oauthController _oauth2Controller.OAuth2Controller
	oauth2Conf      *config.OAuth2
	logger          echo.Logger
}

func (m *authorizingMiddleware) PlayerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauthController.PlayerAuthorizing(pctx, next)
	}
}

func (m *authorizingMiddleware) AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauthController.AdminAuthorizing(pctx, next)
	}
}
