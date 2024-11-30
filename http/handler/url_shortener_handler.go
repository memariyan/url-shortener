package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"url-shortner/http/dto"
	"url-shortner/service"
)

func Convert(c echo.Context) error {
	request, err := getRequestInfo(c)
	if err != nil {
		return err
	}
	result := service.ConvertURL(request.URL)

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
	if err != nil {
		log.Errorln("Failed parse the request body %s", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Debugln("request body : ", request)

	return &request, nil
}
