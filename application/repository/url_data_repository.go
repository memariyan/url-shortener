package repository

import (
	log "github.com/sirupsen/logrus"
	"url-shortner/application/database"
	"url-shortner/application/model"
)

func Save(data *model.URLData) error {
	err := database.DB.Save(data).Error

	if err != nil {
		log.Errorln("error in storing in database: %s", err.Error())
		return err
	}
	return nil
}

func GetByKey(key string) *model.URLData {
	var data model.URLData
	result := database.DB.First(&data, "`key` = ?", key)
	if result != nil {
		return &data
	} else {
		return nil
	}
}

func GetByOriginalUrl(originalUrl string) *model.URLData {
	var data model.URLData
	result := database.DB.First(&data, "`original_url` = ?", originalUrl)
	if result != nil {
		return &data
	} else {
		return nil
	}
}
