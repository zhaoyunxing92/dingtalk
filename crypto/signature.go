package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

//GetSignature 获取签名
func GetSignature(timestamp, secret string, ticket string) string {
	str := timestamp + "\n" + ticket
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape(sign)
}
