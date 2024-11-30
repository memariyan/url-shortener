package service

import "math/rand"

const BaseServerAddress = "http://localhost:8000"

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var storage = make(map[string]string)

func ConvertURL(url string) string {
	pathKey := getNewRandomPathKey()
	storage[pathKey] = url

	return BaseServerAddress + "/" + pathKey
}

func GetOriginalURL(pathKey string) string {
	return storage[pathKey]
}

func getNewRandomPathKey() string {
	newKey := generateRandStringBytes(6)
	if len(storage[newKey]) == 0 {
		return newKey
	}

	return getNewRandomPathKey()
}

func generateRandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}

	return string(b)
}
