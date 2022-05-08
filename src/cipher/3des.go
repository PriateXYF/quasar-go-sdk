package cipher

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// 3DES定义密钥必须是24byte

//解密
func ThriDESDeCrypt(crypted, key []byte) []byte {
	//获取block块
	block, _ := des.NewTripleDESCipher(key)
	//创建切片
	context := make([]byte, len(crypted))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context, crypted)
	//去补码
	context, _ = PKCS7UnPadding(context)
	return context
}

//加密
func ThriDESEnCrypt(origData, key []byte) []byte {
	//获取block块
	block, _ := des.NewTripleDESCipher(key)
	//补码
	origData = PKCS7Padding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	//创建明文长度的数组
	crypted := make([]byte, len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

func EncryptBy3DES(data []byte, key []byte) (string, error) {
	res := ThriDESEnCrypt(data, key)
	return base64.StdEncoding.EncodeToString(res), nil
}

func DecryptBy3DES(data string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return ThriDESDeCrypt(dataByte, key), nil
}

// func main() {
// 	// 定义明文
// 	origData := []byte("测试密钥")
// 	// 加密
// 	en, _ := EncryptBy3DES(origData)
// 	// 解密
// 	de, _ := DecryptBy3DES(en)
// 	fmt.Println(string(de))
// }
