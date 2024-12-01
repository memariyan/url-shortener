package service

import (
	"math/rand"
	"strings"

	"url-shortner/model"
	"url-shortner/repository"
)

const BaseServerAddress = "http://localhost:8001"

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ConvertURL(url string) (string, error) {
	var data *model.URLData

	data = repository.GetByOriginalUrl(url)
	if len(data.Key) != 0 {
		return BaseServerAddress + "/" + data.Key, nil
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

	return BaseServerAddress + "/" + newKey, nil
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
