package helper

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}
func RemoveDomainError(url string) bool {
	newUrl := url
	replaceThis := []string{"http", "https", "www."}
	for i := range replaceThis {
		newUrl = strings.Replace(newUrl, replaceThis[i], "", 1)
	}
	newUrl = strings.Split(newUrl, "/")[0]
	return newUrl != os.Getenv("DOMAIN")
}
