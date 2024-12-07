package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"url-shortner/internal/repository"
	"url-shortner/internal/worker"

	"url-shortner/internal/config"
	"url-shortner/internal/model"
)

type UrlShortenerServiceSuite struct {
	suite.Suite
	repo   *repository.URLDataRepositoryMock
	config *config.Config
}

func (suite *UrlShortenerServiceSuite) SetupTest() {
	suite.config = config.MockConfig()
	suite.repo = repository.MockRepo()
	worker.Get().Start()
}

func TestUrlShortenerService(t *testing.T) {
	suite.Run(t, new(UrlShortenerServiceSuite))
}

func (suite *UrlShortenerServiceSuite) TestUrlShortenerService_ConvertURL_NewURL_Success() {
	require := suite.Require()
	expectedKeyLength := 6

	// given->
	url := "https://www.google.com"
	suite.repo.On("GetByOriginalUrl", url).Return(&model.URLData{})
	suite.repo.On("Save", mock.Anything).Return(nil)

	// when ->
	shortenedUrl, err := ConvertURL(url)

	// then->
	require.NoError(err)
	require.NotNil(shortenedUrl)
	require.True(strings.Contains(shortenedUrl, suite.config.Server.Address))
	key := strings.Split(strings.Replace(shortenedUrl, "http://", "", 1), "/")[1]
	require.True(len(key) == expectedKeyLength)
}

func (suite *UrlShortenerServiceSuite) TestUrlShortenerService_ConvertURL_URLAlreadyExist_Success() {
	require := suite.Require()
	expectedKeyLength := 6
	expectedModel := &model.URLData{OriginalUrl: "https://www.google.com", Key: "JJKoiR", Id: 1}

	// given->
	suite.repo.On("GetByOriginalUrl", expectedModel.OriginalUrl).Return(expectedModel)

	// when ->
	shortenedUrl, err := ConvertURL(expectedModel.OriginalUrl)

	// then->
	require.NoError(err)
	require.NotNil(shortenedUrl)
	require.True(strings.Contains(shortenedUrl, suite.config.Server.Address))
	key := strings.Split(strings.Replace(shortenedUrl, "http://", "", 1), "/")[1]
	require.True(len(key) == expectedKeyLength)
	require.Equal(expectedModel.Key, key)
}
