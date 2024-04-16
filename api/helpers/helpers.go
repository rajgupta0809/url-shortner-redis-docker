package helpers

import "strings"

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == "locahost:3000" {
		return false
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(url, "https://", "", 1)
	newUrl = strings.Replace(url, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	if url == "locahost:3000" {
		return false
	}

	return true
}
