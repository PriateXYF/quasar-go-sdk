package src

import (
	"encoding/base64"
	"fmt"
	"quasar-sdk/src/cipher"
	"strconv"
	"strings"
)

// 加密字符串
func (b business) EncryptoString(str string) (string, error) {
	// 使用
	return b.Encrypto([]byte(str))
}

// 加密数据
func (b business) Encrypto(rawData []byte) (string, error) {
	// 随机生成DEK
	dek := b.GenerateDEK()
	// 使用DEK加密原始数据
	encData, err := cipher.EncryptByAes(rawData, dek)
	if err != nil {
		return "", err
	}
	var encDek string
	var res string
	// 如果处于严格模式下
	if b.IsStrict != 0 {
		// 向密钥管理系统传递DEK，使用密钥管理系统的KEK加密DEK
		// encDek, err = b.EncryptoDEK(dek)
		// if err != nil {
		// 	return "", err
		// }
		// 如果处于非严格模式下
	} else {
		// 从内存中获取当前最新版本KEK
		k := b.GetKek(0)
		if k == nil {
			return "", fmt.Errorf("获取KEK失败!")
		}
		// 使用当前KEK加密DEK
		encDek, err = k.Encrypto(dek)
		if err != nil {
			return "", err
		}
		// 将加密DEK、加密数据、KEK版本号使用.合并
		res, err = b.Combine(encData, encDek, *k)
		if err != nil {
			return "", err
		}
	}

	return res, nil
}

func (b business) GenerateDEK() []byte {
	// 使用
	return cipher.RandBytes(32)
}

// 将加密后的数据、加密后的DEK、密钥版本、业务名称合并
func (b business) Combine(encData string, encDek string, k kek) (string, error) {
	kekVersion := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(k.Version)))
	businessName := base64.StdEncoding.EncodeToString([]byte(b.Name))
	list := []string{encData, encDek, kekVersion, businessName}
	return strings.Join(list, "."), nil
}
