package http

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	MainGroup(e)

	return e
}
