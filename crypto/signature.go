package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
)

//GetSignature 获取签名
func GetSignature(timestamp int64, secret string, ticket string) string {
	str := strconv.FormatInt(timestamp, 10) + "\n" + ticket
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
