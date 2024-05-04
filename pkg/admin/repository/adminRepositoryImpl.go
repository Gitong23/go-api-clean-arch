package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	"github.com/labstack/echo/v4"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) adminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(playerEntity *entities.Player) error {
	return nil
}

func (r *adminRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	return nil, nil
}
