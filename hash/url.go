package hash

import (
	"crypto/rand"
	"encoding/base64"
	"log"
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