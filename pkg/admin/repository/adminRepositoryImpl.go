package repository

import (
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
	_adminException "github.com/Gitong23/go-api-clean-arch/pkg/admin/exception"
	"github.com/labstack/echo/v4"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {

	admin := new(entities.Admin)
	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err != nil {
		r.logger.Errorf("Error creating admin: %v", err.Error())
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {

	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Error finding Admin by ID: %v", err.Error())
		return nil, &_adminException.AdminNotFound{}
	}

	return admin, nil
}
