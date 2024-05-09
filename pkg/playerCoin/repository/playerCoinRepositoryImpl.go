package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepository(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
