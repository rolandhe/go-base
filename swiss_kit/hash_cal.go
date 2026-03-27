package swiss_kit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Hash64Cal(str string) int64 {
	return Hash64CalBytes([]byte(str))
}

func Hash64CalBytes(body []byte) int64 {
	h := CityHash64(body, uint(len(body)))
	return int64(h)
}

func Sha256(data string) string {
	h := sha256.Sum256([]byte(data))

	// 转成十六进制字符串
	hashStr := hex.EncodeToString(h[:])
	return hashStr
}

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
