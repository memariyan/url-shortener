package http

import (
	"github.com/labstack/echo/v4"
	"url-shortner/internal/http/handler"
)

func NewServer() *echo.Echo {
	e := echo.New()
	MainGroup(e)

	return e
}

func MainGroup(e *echo.Echo) {
	e.POST("/convert", handler.Convert)
	e.GET("/:pathKey", handler.Redirect)
}
