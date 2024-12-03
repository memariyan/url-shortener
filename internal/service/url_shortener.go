package service

import (
	"math/rand"
	"strconv"
	"strings"

	"url-shortner/internal/config"
	"url-shortner/internal/model"
	"url-shortner/internal/repository"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var repo repository.URLDataRepository = repository.URLDataRepositoryImpl{}
var configuration = &config.Application

func ConvertURL(url string) (string, error) {
	serverAddress := configuration.Server.Address + ":" + strconv.Itoa(configuration.Server.Port)
	var data *model.URLData

	data = repo.GetByOriginalUrl(url)
	if len(data.Key) != 0 {
		return serverAddress + "/" + data.Key, nil
	}

	newKey := generateRandStringBytes(6)
	data = &model.URLData{OriginalUrl: url, Key: newKey}
	err := repo.Save(data)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ConvertURL(url)
		}
		return "", err
	}

	return serverAddress + "/" + newKey, nil
}

func GetOriginalURL(pathKey string) string {
	data := repo.GetByKey(pathKey)
	if data == nil {
		return ""
	}

	return data.OriginalUrl
}

func generateRandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}

	return string(b)
}
