package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"url-shortner/internal/config"
	mockconfig "url-shortner/mocks/config"

	"url-shortner/internal/http/dto"
	"url-shortner/mocks/database"
)

type UrlShortenerHandlerSuite struct {
	suite.Suite
	client    *http.Client
	server    *httptest.Server
	config    *config.Config
	sqlMock   sqlmock.Sqlmock
	redisMock redismock.ClientMock
}

func TestUrlShortenerHandlers(t *testing.T) {
	suite.Run(t, new(UrlShortenerHandlerSuite))
}

func (suite *UrlShortenerHandlerSuite) SetupTest() {
	suite.server = httptest.NewServer(NewServer())
	suite.client = suite.server.Client()
	suite.sqlMock = database.MockDB()
	suite.redisMock = database.MockRedis()
	suite.config = mockconfig.MockConfig()
}

func (suite *UrlShortenerHandlerSuite) TestURLShortenerHandler_ConvertURL_Success() {
	//given ->
	require := suite.Require()
	url := "https://www.google.com"
	expectedUrlContains := suite.config.Server.Address + ":" + strconv.Itoa(suite.config.Server.Port)

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec("INSERT INTO `url_data` *").
		WithArgs(url, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlMock.ExpectCommit()

	body, _ := json.Marshal(dto.URLShortenerRequest{URL: url})
	request, _ := http.NewRequest(http.MethodPost, suite.server.URL+"/convert", bytes.NewBuffer(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// when->
	response, err := suite.client.Do(request)

	//then->
	require.NoError(err)
	defer response.Body.Close()
	require.Equal(response.StatusCode, http.StatusOK)

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(response.Body)
	responseData := &dto.URLShortenerResponse{}
	err = json.Unmarshal(buf.Bytes(), &responseData)
	require.NoError(err)
	require.NotNil(responseData.Result)
	require.True(strings.Contains(responseData.Result, expectedUrlContains))

	defer suite.server.Close()
}
