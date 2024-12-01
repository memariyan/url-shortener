package repository

import (
	log "github.com/sirupsen/logrus"

	"url-shortner/config"
	"url-shortner/model"
)

func Save(data *model.URLData) error {
	err := config.DB.Save(data).Error

	if err != nil {
		log.Errorln("error in storing in database: %s", err.Error())
		return err
	}
	return nil
}

func GetByKey(key string) *model.URLData {
	var data model.URLData
	result := config.DB.First(&data, "`key` = ?", key)
	if result != nil {
		return &data
	} else {
		return nil
	}
}

func GetByOriginalUrl(originalUrl string) *model.URLData {
	var data model.URLData
	result := config.DB.First(&data, "`original_url` = ?", originalUrl)
	if result != nil {
		return &data
	} else {
		return nil
	}
}
