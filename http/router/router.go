package router

import (
	"github.com/labstack/echo/v4"
	"url-shortner/http/api"
)

func New() *echo.Echo {

	// create a new echo instance
	e := echo.New()

	//set main routes
	api.MainGroup(e)

	return e
}
