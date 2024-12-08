package service

import (
	"strconv"

	gonanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
	"url-shortner/internal/config"
	"url-shortner/internal/model"
	"url-shortner/internal/repository"
	"url-shortner/internal/worker"
)

func ConvertURL(url string) (string, error) {
	serverAddress := config.Get().Server.Address + ":" + strconv.Itoa(config.Get().Server.Port)
	var data *model.URLData

	data = repository.Get().GetByOriginalUrl(url)
	if len(data.Key) != 0 {
		return serverAddress + "/" + data.Key, nil
	}

	if newKey, err := gonanoid.New(6); err != nil {
		return "", err
	} else {
		save(url, newKey)
		return serverAddress + "/" + newKey, nil
	}
}

func GetOriginalURL(pathKey string) string {
	data := repository.Get().GetByKey(pathKey)
	if data == nil {
		return ""
	}

	return data.OriginalUrl
}

func save(url, newKey string) {
	log.Infof("Saving URL: %s with Key: %s", url, newKey)
	data := &model.URLData{OriginalUrl: url, Key: newKey}
	worker.Get().Jobs <- data
}
