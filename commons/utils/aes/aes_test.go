package aes

import (
	"fmt"
	"testing"
)

const KEY = "123456789ABCDEFG"

func TestEncrypt(t *testing.T) {
	rst := Encrypt([]byte("你好，加密世界!"), []byte(KEY))
	fmt.Printf("%x", rst)
}

func TestDecrypt(t *testing.T) {
	rst := Encrypt([]byte("你好，加密世界!"), []byte(KEY))
	rst = Decrypt(rst, []byte(KEY))
	fmt.Printf("%s", rst)
}
