package dingtalk

import (
	"crypto/aes"
	"encoding/base64"
)

const (
	AES_ENCODE_KEY_LENGTH = 43
)

func NewCrypto(token, aesKey, suiteKey string) (c *Crypto) {
	c = &Crypto{
		Token:    token,
		AesKey:   aesKey,
		SuiteKey: suiteKey,
	}
	if len(aesKey) != AES_ENCODE_KEY_LENGTH {
		panic("不合法的aeskey")
	}
	var err error
	c.bkey, err = base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
		panic(err.Error())
	}
	c.block, err = aes.NewCipher(c.bkey)
	if err != nil {
		panic(err.Error())
	}
	return c
}
