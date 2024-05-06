package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	_playerException "github.com/Gitong23/go-api-clean-arch/pkg/player/exception"
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

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {

	player := new(entities.Player)
	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Errorf("Error creating player: %v", err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {

	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Error finding player by ID: %v", err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
