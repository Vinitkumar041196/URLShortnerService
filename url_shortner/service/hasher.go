package service

import (
	"crypto/sha1"
	"encoding/base64"
)

func hashURL(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
