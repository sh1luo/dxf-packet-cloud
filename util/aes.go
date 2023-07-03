package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
)

var key = []byte{84, 72, 73, 82, 73, 83, 77, 90, 83, 69, 67, 82, 69, 84, 75, 69}

func AESCBCEncrypt(src []byte) (string, error) {
	//virtualizersdk.Macro(virtualizersdk.LION_RED_START)

	if len(key) > 16 {
		key = key[:16]
	}
	plaintext := src
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = Padding(plaintext, blockSize)
	if len(plaintext)%aes.BlockSize != 0 {
		log.Println("plaintext is not a multiple of the block size")
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println(err)
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	//virtualizersdk.Macro(virtualizersdk.LION_RED_END)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AESCBCDecrypt(src []byte) (string, error) {
	//virtualizersdk.Macro(virtualizersdk.LION_RED_START)
	ciphertext, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		log.Println(fmt.Errorf("base64 decode err:%s", err))
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(fmt.Errorf("aes.NewCipher err:%s", err))
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		log.Println("ciphertext too short")
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		log.Println("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = UnPadding(ciphertext)
	//virtualizersdk.Macro(virtualizersdk.LION_RED_END)
	return string(ciphertext), nil
}

func Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	char := []byte{byte(padding)}
	newPlain := bytes.Repeat(char, padding)
	return append(plainText, newPlain...)
}

func UnPadding(plainText []byte) []byte {
	length := len(plainText)
	lastChar := plainText[length-1]
	padding := int(lastChar)
	return plainText[:length-padding]
}
