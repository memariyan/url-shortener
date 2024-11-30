package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortner/service"
)

func Redirect(c echo.Context) error {

	pathKey := c.Param("pathKey")
	redirectUrl := service.GetOriginalURL(pathKey)
	if len(redirectUrl) != 0 {
		c.Response().Header().Set("Location", redirectUrl)
		return c.NoContent(http.StatusMovedPermanently)
	} else {
		return c.NoContent(http.StatusNotFound)
	}

}
