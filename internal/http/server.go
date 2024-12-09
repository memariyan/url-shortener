package http

import (
	"github.com/labstack/echo/v4"
	"url-shortner/internal/http/handler"
	"url-shortner/internal/http/middleware"
)

func NewHttpServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.TracingMiddleware)

	MainGroup(e)
	MetricGroup(e)

	return e
}

func MetricGroup(e *echo.Echo) {
	e.GET("/metrics", handler.Metric())
}

func MainGroup(e *echo.Echo) {
	e.POST("/convert", handler.Convert)
	e.GET("/:pathKey", handler.Redirect)
}
