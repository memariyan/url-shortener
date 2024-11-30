package http

import (
	"github.com/labstack/echo/v4"

	"url-shortner/http/handler"
)

func MainGroup(e *echo.Echo) {
	e.POST("/convert", handler.Convert)
	e.GET("/:pathKey", handler.Redirect)
}
