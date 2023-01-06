package pkcs7

import (
	"bytes"
	"errors"
)

func Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func UnPad(src []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid blockSize")
	}

	if len(src) == 0 {
		return nil, errors.New("unPad error")
	}
	c := src[len(src)-1]
	n := int(c)

	if n == 0 || n > len(src) {
		return nil, errors.New("unPad error")
	}

	for i := 0; i < n; i++ {
		if src[len(src)-n+i] != c {
			return nil, errors.New("unPad error")
		}
	}
	return src[:len(src)-n], nil
}
