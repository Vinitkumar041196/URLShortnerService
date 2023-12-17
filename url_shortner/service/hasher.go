package service

import (
	"crypto/sha1"
	"encoding/base64"
)

//generate the hash of URL and then convert it to base64 url encoded string 
func hashURL(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	encodedURL := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encodedURL
}
