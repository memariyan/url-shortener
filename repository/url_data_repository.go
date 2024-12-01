package repository

import (
	log "github.com/sirupsen/logrus"

	"url-shortner/config"
	"url-shortner/model"
)

var db = config.DatabaseConnection()

func Save(data *model.URLData) error {
	err := db.Save(data).Error

	if err != nil {
		log.Errorln("error in storing in database: %s", err.Error())
		return err
	}
	return nil
}

func GetByKey(key string) *model.URLData {
	var data model.URLData
	result := db.First(&data, "`key` = ?", key)
	if result != nil {
		return &data
	} else {
		return nil
	}
}

func GetByOriginalUrl(originalUrl string) *model.URLData {
	var data model.URLData
	result := db.First(&data, "`original_url` = ?", originalUrl)
	if result != nil {
		return &data
	} else {
		return nil
	}
}
