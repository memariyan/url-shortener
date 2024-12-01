package service

import (
	"math/rand"
	"strconv"
	"strings"

	"url-shortner/application/config"
	"url-shortner/application/model"
	"url-shortner/application/repository"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ConvertURL(url string) (string, error) {
	serverAddress := config.Application.Server.Address + ":" + strconv.Itoa(config.Application.Server.Port)
	var data *model.URLData

	data = repository.GetByOriginalUrl(url)
	if len(data.Key) != 0 {
		return serverAddress + "/" + data.Key, nil
	}

	newKey := generateRandStringBytes(6)
	data = &model.URLData{OriginalUrl: url, Key: newKey}
	err := repository.Save(data)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ConvertURL(url)
		}
		return "", err
	}

	return serverAddress + "/" + newKey, nil
}

func GetOriginalURL(pathKey string) string {
	data := repository.GetByKey(pathKey)
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
