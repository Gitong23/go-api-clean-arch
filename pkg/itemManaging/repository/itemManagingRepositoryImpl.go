package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemManagingRepositoryImp struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepository(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImp {
	return &itemManagingRepositoryImp{
		db:     db,
		logger: logger,
	}
}
