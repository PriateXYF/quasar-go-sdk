package cipher

import (
	"crypto/md5"
	"fmt"
)

// md5加密字符串
func MD5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 三次md5加密字符串
func ThriMD5String(str string) string {
	return MD5String(MD5String(MD5String(str)))
}
