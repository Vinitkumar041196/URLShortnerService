package service

import (
	"net/url"
	"strings"
)

// used to get the domain part of URL
func GetDomainFromURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	parts := strings.Split(u.Hostname(), ".")
	if len(parts) > 1 {
		return parts[len(parts)-2] + "." + parts[len(parts)-1]
	}
	return ""
}
