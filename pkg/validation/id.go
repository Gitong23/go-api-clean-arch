package validation

import (
	_admin "github.com/Gitong23/go-api-clean-arch/pkg/admin/exception"
	_player "github.com/Gitong23/go-api-clean-arch/pkg/player/exception"
	"github.com/labstack/echo/v4"
)

func AdminIDGetting(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_admin.AdminNotFound{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}

func PlayerIDGetting(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_player.PlayerNotFound{PlayerID: "Unknown"}
	} else {
		return playerID, nil
	}
}
