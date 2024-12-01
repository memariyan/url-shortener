package http

import (
	"github.com/labstack/echo/v4"
	handler2 "url-shortner/application/http/handler"
)

func MainGroup(e *echo.Echo) {
	e.POST("/convert", handler2.Convert)
	e.GET("/:pathKey", handler2.Redirect)
}
