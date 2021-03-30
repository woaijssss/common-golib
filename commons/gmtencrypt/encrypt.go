package gmtencrypt

import (
	"encoding/hex"
	"fmt"
	"github.com/woaijssss/common-golib/commons/constant"
	"github.com/woaijssss/common-golib/commons/utils/aes"
)

// 手机号加密
func Encrypt(raw string) string {
	return fmt.Sprintf("%x", aes.Encrypt([]byte(raw), []byte(constant.AES_PRIVATE_KEY)))
}

// 手机号解密
func Decrypt(raw string) string {
	bytes, err := hex.DecodeString(raw)
	if err != nil {
		return ""
	}

	return string(aes.Decrypt(bytes, []byte(constant.AES_PRIVATE_KEY)))
}
