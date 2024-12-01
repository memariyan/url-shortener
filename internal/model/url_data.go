package model

import "encoding/json"

type URLData struct {
	Id int `gorm:"type:int;primary_key" json:"id"`

	OriginalUrl string `gorm:"type:varchar(255)" json:"original_url"`

	Key string `gorm:"type:varchar(10)" json:"key"`
}

func (s URLData) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (data *URLData) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, data)
}
