package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
	MySQL  MySQL  `mapstructure:"mysql"`
}

type Server struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

var ApplicationConfig Config

func ReadConfig() *Config {
	ApplicationConfig = Config{}
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&ApplicationConfig)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &ApplicationConfig
}
