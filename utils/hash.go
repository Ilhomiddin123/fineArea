package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	secretKey = "serzxd+ct65123fvbhnjmlkjnilbuvcyfxtgcjgbnkmkjbhvcfxgdg++wkl3op2rjknsdstchj"
)

func GenerateMD5Hash(requestData string) string {
	data := []byte(requestData + secretKey)

	// Генерация MD5 хэша
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
