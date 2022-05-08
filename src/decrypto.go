package src

import (
	"encoding/base64"
	"fmt"
	"github.com/PriateXYF/quasar-go-sdk/src/cipher"
	"strconv"
	"strings"
)

// 解密字符串
func (b business) DecryptoString(str string) (string, error) {
	// 使用
	rawData, err := b.Decrypto(str)
	if err != nil {
		return "", err
	}
	return string(rawData), nil
}

// 解密数据
func (b business) Decrypto(encData string) ([]byte, error) {
	// 解析加密数据
	data, err := b.ParseData(encData)
	if err != nil {
		return nil, err
	}
	var dek []byte
	// 如果处于严格模式下
	if b.IsStrict != 0 {
		// 向密钥管理系统传递加密DEK，使用密钥管理系统的KEK解密DEK
		// dek, err = b.DecryptoDEK(data.EncDek)
		// if err != nil {
		// 	return nil, err
		// }
		// 如果处于非严格模式下
	} else {
		// 从内存中获取对应版本KEK
		kek := b.GetKek(data.KekVersion)
		if kek == nil {
			return nil, fmt.Errorf("KEK不存在!")
		}
		// 使用KEK解密DEK
		dek, err = kek.Decrypto(data.EncDek)
		if err != nil {
			return nil, err
		}
	}
	// 使用dek解密加密数据
	rawData, err := cipher.DecryptByAes(data.EncData, dek)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func (b business) ParseData(encStr string) (*EncryptoData, error) {
	list := strings.Split(encStr, ".")
	if len(list) != 4 {
		return nil, fmt.Errorf("解析密文失败!字符切割异常!")
	}
	res := new(EncryptoData)
	res.EncData = list[0]
	res.EncDek = list[1]
	kekVersion, err := base64.StdEncoding.DecodeString(list[2])
	if err != nil {
		return nil, err
	}
	businessName, err := base64.StdEncoding.DecodeString(list[3])
	if err != nil {
		return nil, err
	}
	res.KekVersion, err = strconv.Atoi(string(kekVersion))
	if err != nil {
		return nil, err
	}
	res.BusinessName = string(businessName)
	return res, nil

}
