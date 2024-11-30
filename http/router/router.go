package router

import (
	"github.com/labstack/echo/v4"

	"url-shortner/http/api"
)

func New() *echo.Echo {
	e := echo.New()
	api.MainGroup(e)

	return e
}
