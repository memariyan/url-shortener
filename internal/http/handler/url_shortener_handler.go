package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"url-shortner/internal/http/dto"
	"url-shortner/internal/service"
)

func Convert(c echo.Context) error {
	request, err := getRequestInfo(c)
	if err != nil {
		log.Errorln("Failed parse the request body %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := service.ConvertURL(request.URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.URLShortenerResponse{
		Result: result,
	})
}

func getRequestInfo(c echo.Context) (*dto.URLShortenerRequest, error) {
	request := dto.URLShortenerRequest{}
	defer c.Request().Body.Close()
	err := c.Bind(&request)
	if err == nil {
		err = request.Validate()
	}

	return &request, err
}
