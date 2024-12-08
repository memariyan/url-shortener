package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func TracingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract context from the incoming request
		ctx := c.Request().Context()
		propagator := otel.GetTextMapPropagator()
		ctx = propagator.Extract(ctx, propagation.HeaderCarrier(c.Request().Header))

		// Start a new span
		tracer := otel.Tracer("url-shortener")
		spanCtx, span := tracer.Start(ctx, c.Request().Method+" "+c.Request().URL.Path, trace.WithAttributes(
			semconv.HTTPMethodKey.String(c.Request().Method),
			semconv.HTTPTargetKey.String(c.Request().URL.Path),
			semconv.HTTPClientIPKey.String(c.RealIP()),
		))
		defer span.End()

		// Inject span context back into the request
		c.SetRequest(c.Request().WithContext(spanCtx))

		// Call the next handler
		err := next(c)

		// Set additional span attributes based on response
		statusCode := c.Response().Status
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(statusCode))
		if statusCode >= http.StatusInternalServerError {
			span.SetStatus(500, "Internal server error")
		} else {
			span.SetStatus(200, "OK")
		}

		return err
	}
}
