package config

func MockConfig() *Config {
	configuration = Config{Server: Server{Address: "http://localhost", Port: 8000}}
	return &configuration
}
