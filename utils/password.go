package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetPasswordHash(password, key string) string {
	original := password + "_$_" + key
	hash := md5.Sum([]byte(original))
	return hex.EncodeToString(hash[:])
}
