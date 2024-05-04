package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/labstack/echo/v4"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) PlayerRepository {
	return &playerRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
