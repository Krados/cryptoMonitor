package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func HmacSha256(secret, payload string) (signature string) {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func NowInMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
