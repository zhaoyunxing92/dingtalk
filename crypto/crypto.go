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
	//"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

const (
	AesEncodeKeyLength = 43
	alphabets          = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type Encrypt struct {
	Text string `json:"encrypt"`
}

type DingTalkCrypto struct {
	Token          string
	EncodingAESKey string
	SuiteKey       string
	BKey           []byte
	Block          cipher.Block
}

func (c *DingTalkCrypto) GetDecryptMsg(signature, timestamp, nonce, secretMsg string) (string, error) {
	if !c.VerificationSignature(c.Token, timestamp, nonce, secretMsg, signature) {
		return "", errors.New("ERROR: 签名不匹配")
	}
	decode, err := base64.StdEncoding.DecodeString(secretMsg)
	if err != nil {
		return "", err
	}
	if len(decode) < aes.BlockSize {
		return "", errors.New("ERROR: 密文太短")
	}
	blockMode := cipher.NewCBCDecrypter(c.Block, c.BKey[:c.Block.BlockSize()])
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = pkCS7UnPadding(plantText)
	size := binary.BigEndian.Uint32(plantText[16:20])
	plantText = plantText[20:]
	corpID := plantText[size:]
	if string(corpID) != c.SuiteKey {
		return "", errors.New("ERROR: CorpID匹配不正确")
	}
	return string(plantText[:size]), nil
}

func (c *DingTalkCrypto) GetEncryptMsg(msg string) (map[string]string, error) {
	var timestamp = time.Now().Second()
	var nonce = c.RandomString(12)
	str, sign, err := c.GetEncryptMsgDetail(msg, fmt.Sprint(timestamp), nonce)

	return map[string]string{"nonce": nonce, "timeStamp": fmt.Sprint(timestamp), "encrypt": str, "msg_signature": sign}, err
}
func (c *DingTalkCrypto) GetEncryptMsgDetail(msg, timestamp, nonce string) (string, string, error) {
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(msg)))
	msg = c.RandomString(16) + string(size) + msg + c.SuiteKey
	plantText := pkCS7Padding([]byte(msg), c.Block.BlockSize())
	if len(plantText)%aes.BlockSize != 0 {
		return "", "", errors.New("ERROR: 消息体size不为16的倍数")
	}
	blockMode := cipher.NewCBCEncrypter(c.Block, c.BKey[:c.Block.BlockSize()])
	chipherText := make([]byte, len(plantText))
	blockMode.CryptBlocks(chipherText, plantText)
	outMsg := base64.StdEncoding.EncodeToString(chipherText)
	signature := c.CreateSignature(c.Token, timestamp, nonce, string(outMsg))
	return string(outMsg), signature, nil
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
func (c *DingTalkCrypto) CreateSignature(token, timestamp, nonce, msg string) string {
	params := make([]string, 0)
	params = append(params, token)
	params = append(params, timestamp)
	params = append(params, nonce)
	params = append(params, msg)
	sort.Strings(params)
	return sha1Sign(strings.Join(params, ""))
}

// VerificationSignature 验证数据签名
func (c *DingTalkCrypto) VerificationSignature(token, timestamp, nonce, msg, sigture string) bool {
	return c.CreateSignature(token, timestamp, nonce, msg) == sigture
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
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
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
