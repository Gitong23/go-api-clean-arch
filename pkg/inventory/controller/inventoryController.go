package controller

import "github.com/labstack/echo/v4"

type InventroryController interface {
	Listing(pctx echo.Context) error
}
