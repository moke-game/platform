package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// CBCDecrypt AES-CBC 解密
func CBCDecrypt(key []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	if len(ciphercode)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphercode length is not a multiple of the block size")
	}

	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)
	ciphercode = bytes.Trim(ciphercode, "\x00")
	return strings.TrimSpace(string(ciphercode)), nil
}

// EncryptMD5 encrypt data with md5.
func EncryptMD5(data, key string) (string, error) {
	str := fmt.Sprintf("%s.%s", data, key)
	_, err := md5.New().Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5.New().Sum(nil)), nil
}
