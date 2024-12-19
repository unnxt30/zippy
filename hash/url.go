package hash

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/url"
)

// GenerateRandomString generates a random string of the specified length
func GetShortURL(n int) string {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if (err != nil) {
        log.Fatal(err)
    }
    return base64.URLEncoding.EncodeToString(b)[:n]
}

func ValidURL(testURL string) bool {
	parsedURL, err := url.ParseRequestURI(testURL)
	if err != nil {
		return false
	}
	// Ensure the scheme and host are not empty
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}