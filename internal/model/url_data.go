package model

type URLData struct {
	Id int `gorm:"type:int;primary_key"`

	OriginalUrl string `gorm:"type:varchar(255)"`

	Key string `gorm:"type:varchar(10)"`
}
