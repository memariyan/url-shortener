package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configuration *Config

func App() *Config {
	if configuration == nil {
		configuration = readConfig()
	}
	return configuration
}

type Config struct {
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Worker Worker `yaml:"worker"`
	Jaeger Jaeger `yaml:"jaeger"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Worker struct {
	Size int `yaml:"size"`
}

func readConfig() *Config {
	configuration = &Config{}
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return configuration
}
