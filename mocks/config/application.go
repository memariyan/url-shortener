package config

import "url-shortner/internal/config"

func MockConfig() *config.Config {
	config.Application = config.Config{Server: config.Server{Address: "http://localhost", Port: 8000}}
	return &config.Application
}
