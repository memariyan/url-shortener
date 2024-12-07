package repository

import (
	"context"

	log "github.com/sirupsen/logrus"

	"url-shortner/internal/database"
	"url-shortner/internal/model"
)

type URLDataRepository interface {
	Save(data *model.URLData) error

	GetByKey(key string) *model.URLData

	GetByOriginalUrl(originalUrl string) *model.URLData
}

type URLDataRepositoryImpl struct {
}

var instance URLDataRepository = &URLDataRepositoryImpl{}
var ctx = context.Background()

func Get() URLDataRepository {
	return instance
}

func (r URLDataRepositoryImpl) Save(data *model.URLData) error {
	err := database.GetDB().Save(data).Error
	if err != nil {
		log.Errorln("error in storing in database:", err.Error())
		return err
	}

	return nil
}

func (r URLDataRepositoryImpl) GetByKey(key string) *model.URLData {
	var data *model.URLData
	if data = readFromCache(key); data != nil {
		return data
	}
	if result := database.GetDB().First(&data, "`key` = ?", key); result != nil {
		writeOnCache(key, data)
		return data
	} else {
		return nil
	}
}

func readFromCache(key string) *model.URLData {
	var result model.URLData
	bytes, _ := database.GetRedis().Get(ctx, key).Bytes()
	if err := result.UnmarshalBinary(bytes); err == nil {
		return &result
	}

	return nil
}

func writeOnCache(key string, data *model.URLData) {
	if err := database.GetRedis().Set(ctx, key, data, 0).Err(); err != nil {
		log.Errorln(err)
	}
}

func (r URLDataRepositoryImpl) GetByOriginalUrl(originalUrl string) *model.URLData {
	var data model.URLData
	if result := database.GetDB().First(&data, "`original_url` = ?", originalUrl); result != nil {
		return &data
	}

	return nil
}
