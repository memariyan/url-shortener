package repository

import (
	"github.com/stretchr/testify/mock"

	"url-shortner/internal/model"
)

type URLDataRepositoryMock struct {
	mock.Mock
}

func (m *URLDataRepositoryMock) Save(data *model.URLData) error {
	return m.Called(data).Error(0)
}

func (m *URLDataRepositoryMock) GetByKey(key string) *model.URLData {
	return m.Called(key).Get(0).(*model.URLData)
}

func (m *URLDataRepositoryMock) GetByOriginalUrl(originalUrl string) *model.URLData {
	return m.Called(originalUrl).Get(0).(*model.URLData)
}

func MockRepo() *URLDataRepositoryMock {
	mockRepo := &URLDataRepositoryMock{}
	instance = mockRepo
	return mockRepo
}
