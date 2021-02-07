package gmtencrypt

import (
	"encoding/hex"
	"fmt"
	"github.com/woaijssss/common-golib/commons/constant"
	"github.com/woaijssss/common-golib/commons/utils"
	"github.com/woaijssss/common-golib/commons/utils/aes"
)

func MobileLst4(mobile string) string {
	if len(mobile) < 4 {
		return mobile
	} else {
		return mobile[len(mobile)-4:]
	}
}

func MobileMask(mobile string) string {
	if len(mobile) >= 7 {
		return mobile[0:3] + "****" + mobile[7:]
	} else {
		return mobile
	}
}

func MobileMaskPtr(mobile string) *string {
	mask := MobileMask(mobile)
	return util.StringPtr(mask)
}

func CardIDMask(id string) string {
	if len(id) >= 14 {
		return id[0:6] + "********" + id[14:]
	} else {
		return id
	}
}

func CardIDMaskPtr(id string) *string {
	return util.StringPtr(CardIDMask(id))
}

//住户手机号加密
func Encrypt(raw string) string {
	return fmt.Sprintf("%x", aes.Encrypt([]byte(raw), []byte(constant.AES_PRIVATE_KEY)))
}

//住户手机号解密
func Decrypt(raw string) string {
	bytes, err := hex.DecodeString(raw)
	if err != nil {
		return ""
	}

	return string(aes.Decrypt(bytes, []byte(constant.AES_PRIVATE_KEY)))
}
