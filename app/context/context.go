package context

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/woaijssss/common-golib/common"
	"strings"
)

const (
	Token     = "token"
	Name      = "name"
	Email     = "email"
	ErrorCode = "error_code"
	RequestId = "request_id"
)

func GetGinContextWithRequestId() *gin.Context {
	ctx := gin.Context{}

	value := uuid.Must(uuid.NewV4(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	requestId := hex.EncodeToString(m.Sum(nil))
	SetRequestID(&ctx, requestId)
	return &ctx
}

func SetContextData(ctx *gin.Context, key string, value interface{}) {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys[key] = value
}

func GetContextData(ctx *gin.Context, key string) interface{} {
	if ctx.Keys == nil {
		return nil
	}
	value, ok := ctx.Keys[key]
	if ok == false {
		return nil
	}
	return value
}

func SetUserName(ctx *gin.Context, name string) {
	SetContextData(ctx, Name, name)
}

func GetUserName(ctx *gin.Context) string {
	v := GetContextData(ctx, Name)
	if v == nil {
		return ""
	}
	return v.(string)
}

func SetToken(ctx *gin.Context, token string) {
	SetContextData(ctx, Token, token)
}

func GetToken(ctx *gin.Context) string {
	token := GetContextData(ctx, Token)
	return token.(string)
}

func SetErrorCode(ctx *gin.Context, errorCode int) {
	SetContextData(ctx, ErrorCode, errorCode)
}

func GetErrorCode(ctx *gin.Context) int32 {
	v := GetContextData(ctx, ErrorCode)
	if v == nil {
		return common.SUCCESS
	}
	return v.(int32)
}

func SetRequestID(ctx *gin.Context, requestId string) {
	SetContextData(ctx, RequestId, requestId)
}

func GetRequestID(ctx *gin.Context) string {
	v := GetContextData(ctx, RequestId)
	if v == nil {
		return ""
	}
	return v.(string)
}

func GetRequestHost(ctx *gin.Context) (host string) {
	strHost := ctx.Request.Host
	arrHost := strings.Split(strHost, ":")
	if len(arrHost) > 0 {
		host = arrHost[0]
	}
	return
}

func GetRemoteIp(ctx *gin.Context) (host string) {
	strHost := ctx.Request.RemoteAddr
	arrHost := strings.Split(strHost, ":")
	if len(arrHost) == 0 {
		host = "127.0.0.1"
	} else {
		host = arrHost[0]
	}
	return
}
