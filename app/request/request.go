package request

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/woaijssss/common-golib/app/context"
	"github.com/woaijssss/common-golib/app/logger"
	"github.com/woaijssss/common-golib/common"
	"io/ioutil"
	"net/http"
	"time"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code      int32       `json:"code"`
	Msg       string      `json:"msg"`
	RequestId string      `json:"request_id"`
	Time      time.Time   `json:"time"`
	Data      interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int32 , data interface{}) {
	context.SetErrorCode(g.C, errCode)
	response := Response{
		Code:      errCode,
		Data:      data,
		Time:      time.Now(),
		RequestId: context.GetRequestID(g.C),
	}
	res, err := json.Marshal(response)
	if err != nil {
		logger.Errorf(g.C, "smart_request_out: marshal error [%s] ", err.Error())
	}
	logger.Infof(g.C, "smart_request_out: res[%s] ", string(res))
	g.C.JSON(int(httpCode), response)
	return
}

func (g *Gin) ResponseWithMessage(httpCode, errCode int32 , data interface{}, msg string) {
	context.SetErrorCode(g.C, errCode)
	response := Response{
		Code:      errCode,
		Msg: 		msg,
		Data:      data,
		Time:      time.Now(),
		RequestId: context.GetRequestID(g.C),
	}
	res, err := json.Marshal(response)
	if err != nil {
		logger.Errorf(g.C, "smart_request_out: marshal error [%s] ", err.Error())
	}
	logger.Infof(g.C, "smart_request_out: res[%s] ", string(res))
	g.C.JSON(int(httpCode), response)
	return
}

func (g *Gin) ResponseSuccess(data interface{}) {
	g.Response(http.StatusOK, common.SUCCESS, data)
}

func (g *Gin) ResponseError(errCode int32, data interface{}) {
	g.Response(http.StatusOK, errCode,  data)
}

func (g *Gin) ResponseErrorWithMsg(errCode int32, msg string) {
	g.ResponseWithMessage(http.StatusOK, errCode,  nil, msg)
}

func (g *Gin) ResponseErrorWithMsgAndData(errCode int32, data interface{}, msg string) {
	g.ResponseWithMessage(http.StatusOK, errCode,  data, msg)
}

func (g *Gin) BindAndValid(form interface{}) error {
	err := g.C.ShouldBind(form)
	if err != nil {
		logger.Warn(g.C, err.Error())
		return err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logger.Warn(g.C, err.Error())
		return err
	}
	if !check {
		logger.Warn(g.C, "valid check failed")
		return valid.Errors[0]
	}

	return nil
}

func GetRequestId() (requestId string) {
	value := uuid.Must(uuid.NewV4(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func AddRequestId(ctx *gin.Context) {
	requestId := GetRequestId()
	context.SetRequestID(ctx, requestId)
	body, err := ctx.GetRawData()
	if err != nil {
		logger.Error(ctx, err.Error())
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//输入参数写入日志
	logger.Infof(ctx, "smart_request_in: RequestURI[%s],RequestBody[%v]", ctx.Request.Header, string(body))
}
