package logger

import (
	"fmt"
	"github.com/woaijssss/common-golib/app/context"
	"testing"
)

func TestLogger(t *testing.T) {
	//获取日志的存储路径
	filePath := fmt.Sprintf("%s%s",
		"runtime/",
		"logs/",
	)
	fileName := fmt.Sprintf("%s.%s",
		"parking",
		"log",
	)
	//获取配置的日志级别
	logLevel := "debug"
	runMode := "debug"
	expireDay := 1
	//format := "%Y%m%d"
	Setup(filePath, fileName, logLevel, runMode, int32(expireDay))

	ctx := context.GetGinContextWithRequestId()
	Info(ctx, "OOOO")
}
