package src

import (
	"fmt"
	"github.com/PriateXYF/quasar-go-sdk/src/cipher"
	"strings"
)

// 加密数据
func (k kek) Encrypto(data []byte) (string, error) {
	// 针对不同算法进行加密
	switch strings.ToLower(k.Algorithm) {
	case "aes":
		return cipher.EncryptByAes(data, []byte(k.Content))
	case "3des":
		return cipher.EncryptBy3DES(data, []byte(k.Content))
	case "chanchan20":
		return cipher.EncryptByChanChan20(string(data), k.Content)
	}
	return "", fmt.Errorf("算法 %s 不存在", k.Algorithm)
}

// 加密数据
func (k kek) Decrypto(data string) ([]byte, error) {
	// 针对不同算法进行加密
	switch strings.ToLower(k.Algorithm) {
	case "aes":
		return cipher.DecryptByAes(data, []byte(k.Content))
	case "3des":
		return cipher.DecryptBy3DES(data, []byte(k.Content))
	case "chanchan20":
		return cipher.DecryptByChanChan20([]byte(data), k.Content)
	}
	return nil, fmt.Errorf("算法 %s 不存在", k.Algorithm)
}
