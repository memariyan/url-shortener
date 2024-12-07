package config

func MockConfig() *Config {
	configuration = &Config{
		Server: Server{Address: "http://localhost", Port: 8000},
		Worker: Worker{Size: 10}}

	return configuration
}
