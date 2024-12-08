package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Metric() echo.HandlerFunc {
	prometheusHandler := promhttp.Handler()
	return func(ctx echo.Context) error {
		prometheusHandler.ServeHTTP(ctx.Response(), ctx.Request())

		return nil
	}
}
