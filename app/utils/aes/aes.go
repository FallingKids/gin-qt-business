package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AES struct {
	Key []byte
	Iv  []byte
}

func (a *AES) Encrypt(data []byte) (string, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	// PKCS7Padding
	data = pkcs7Padding(data, block.BlockSize())

	ciphertext := make([]byte, len(data))

	// CBC mode
	mode := cipher.NewCBCEncrypter(block, a.Iv)
	mode.CryptBlocks(ciphertext, data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (a *AES) Decrypt(data string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))

	// CBC mode
	mode := cipher.NewCBCDecrypter(block, a.Iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// PKCS7Padding
	plaintext = pkcs7UnPadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding 填充函数
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7Padding 去除函数
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
