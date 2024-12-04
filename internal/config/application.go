package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configuration Config

func Get() *Config {
	return &configuration
}

type Config struct {
	Server Server `mapstructure:"server"`
	MySQL  MySQL  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
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

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func ReadConfig() *Config {
	configuration = Config{}
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &configuration
}
