package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"

	"url-shortner/http/dto"
	"url-shortner/service"
)

func Convert(c echo.Context) error {
	request, err := parseRequestAndValidate(c)
	if err != nil {
		return err
	}
	result := service.ConvertURL(request.URL)

	return c.JSON(http.StatusOK, dto.URLShortenerResponse{
		Result: result,
	})
}

func parseRequestAndValidate(c echo.Context) (*dto.URLShortenerRequest, error) {

	//todo : binding
	request := dto.URLShortenerRequest{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		log.Debugln("Failed reading the request body %s", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed reading the request body!")
	}

	if len(request.URL) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "URL is required")
	}
	log.Debugln("request body : ", request)
	return &request, nil
}
