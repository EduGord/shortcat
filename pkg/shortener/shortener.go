package shortener

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortURL() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
