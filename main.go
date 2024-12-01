package main

import (
	"strconv"

	"url-shortner/config"
	"url-shortner/http"
)

func main() {
	// configuration
	config.ReadConfig()

	// database
	config.DatabaseConnection(&config.ApplicationConfig.MySQL)

	// http server
	e := http.New()
	err := e.Start(":" + strconv.Itoa(config.ApplicationConfig.Server.Port))
	e.Logger.Fatal(err)
}
