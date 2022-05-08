package cipher

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"strings"

	"golang.org/x/crypto/chacha20"
)

type ChaCha20Reader struct {
	baseReader io.Reader
	passByte   [32]byte
	cipher     *chacha20.Cipher
}

func NewChaCha20Reader(pass string, reader io.Reader) *ChaCha20Reader {
	c := new(ChaCha20Reader)
	c.baseReader = reader
	c.passByte = sha256.Sum256([]byte(pass))
	nonce := make([]byte, 24)
	c.cipher, _ = chacha20.NewUnauthenticatedCipher(c.passByte[:32], nonce)
	return c
}

func (c *ChaCha20Reader) Read(p []byte) (n int, err error) {
	n, err = c.baseReader.Read(p)
	if err != nil {
		return n, err
	}
	p = p[:n]
	c.cipher.XORKeyStream(p, p)
	return n, nil
}

func EncryptByChanChan20(msg string, key string) (string, error) {
	r := NewChaCha20Reader(key, strings.NewReader(msg))
	var enc bytes.Buffer
	for {
		ret := make([]byte, 4, 4)
		n, err := r.Read(ret)
		if err != nil {
			break
		}
		enc.Write(ret[:n])
	}
	return base64.StdEncoding.EncodeToString(enc.Bytes()), nil
}

func DecryptByChanChan20(enc []byte, key string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(string(enc))
	if err != nil {
		return nil, err
	}
	// 解密
	r2 := NewChaCha20Reader(key, bytes.NewReader(dataByte))
	var dec bytes.Buffer
	for {
		ret := make([]byte, 4, 4)
		n, err := r2.Read(ret)
		if err != nil {
			break
		}
		dec.Write(ret[:n])
	}
	return dec.Bytes(), nil
}

// func main() {
// 	enc, _ := EncryptByChanChan20("测试内容", "1234")
// 	println(enc)
// 	dec, _ := DecryptByChanChan20([]byte(enc), "1234")
// 	println(string(dec))
// }
