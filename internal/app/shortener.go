package app

import (
	"bytes"
	"encoding/base64"
)

func GenerateShortLink(link string) string {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.RawURLEncoding, &buf)
	enc.Write([]byte(link))
	str := buf.String()
	return str[:len(str)/2]
}
