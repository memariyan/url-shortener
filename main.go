package main

import (
	"url-shortner/http"
)

func main() {
	e := http.New()
	e.Logger.Fatal(e.Start(":8000"))
}
