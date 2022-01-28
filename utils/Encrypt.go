package utils

import (
	"crypto/md5"
	"fmt"
)

func EncryptPassword(password string) string {
	data := []byte(password)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
