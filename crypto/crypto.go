/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"strconv"

	//"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	AesEncodeKeyLength = 43
	alphabets          = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type DingTalkCrypto struct {

	// 签名 token
	Token string

	// 小程序的key
	SuiteKey string

	// 加密 aes_key
	AESKey []byte

	Block cipher.Block
}

// Decrypt 解密
func (c *DingTalkCrypto) Decrypt(encrypt, sign, timestamp, nonce string) (string, error) {
	if !c.VerificationSignature(encrypt, sign, timestamp, nonce) {
		return "", errors.New("签名不匹配")
	}
	decode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	if len(decode) < aes.BlockSize {
		return "", errors.New("密文太短")
	}
	blockMode := cipher.NewCBCDecrypter(c.Block, c.AESKey[:c.Block.BlockSize()])
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = pkCS7UnPadding(plantText)
	size := binary.BigEndian.Uint32(plantText[16:20])
	plantText = plantText[20:]
	corpId := plantText[size:]
	if string(corpId) != c.SuiteKey {
		return "", errors.New("SuiteKey匹配不正确")
	}
	return string(plantText[:size]), nil
}

// Encrypt 加密
func (c *DingTalkCrypto) Encrypt(msg string) (*DingTalkEncrypt, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	nonce := c.RandomString(12)

	encrypt, sign, err := c.GetEncryptMsgDetail(msg, timestamp, nonce)
	return NewDingTalkEncrypt(encrypt, sign, timestamp, nonce), err
}

func (c *DingTalkCrypto) GetEncryptMsgDetail(msg, timestamp, nonce string) (string, string, error) {
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(msg)))
	msg = c.RandomString(16) + string(size) + msg + c.SuiteKey
	plantText := pkCS7Padding([]byte(msg), c.Block.BlockSize())
	if len(plantText)%aes.BlockSize != 0 {
		return "", "", errors.New("消息体size不为16的倍数")
	}
	blockMode := cipher.NewCBCEncrypter(c.Block, c.AESKey[:c.Block.BlockSize()])
	chipherText := make([]byte, len(plantText))
	blockMode.CryptBlocks(chipherText, plantText)
	encrypt := base64.StdEncoding.EncodeToString(chipherText)
	signature := c.CreateSignature(c.Token, timestamp, nonce, encrypt)
	return encrypt, signature, nil
}

func sha1Sign(s string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", bs)
}

// CreateSignature 数据签名
func (c *DingTalkCrypto) CreateSignature(token, encrypt, timestamp, nonce string) string {
	params := make([]string, 0)
	params = append(params, token)
	params = append(params, timestamp)
	params = append(params, nonce)
	params = append(params, encrypt)
	sort.Strings(params)
	return sha1Sign(strings.Join(params, ""))
}

// VerificationSignature 验证数据签名
func (c *DingTalkCrypto) VerificationSignature(encrypt, sign, timestamp, nonce string) bool {
	return c.CreateSignature(c.Token, encrypt, timestamp, nonce) == sign
}

// pkCS7UnPadding 解密补位
func pkCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// pkCS7Padding 加密补位
func pkCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, text...)
}

// RandomString 随机字符串
func (c *DingTalkCrypto) RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var str []byte
	for i := 0; i < length; i++ {
		str = append(str, alphabets[rand.Intn(len(alphabets))])
	}
	return string(str)
}
