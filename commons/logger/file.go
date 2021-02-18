package logger

import (
	"fmt"
	"github.com/woaijssss/common-golib/commons/setting"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s.%s.%s",
		setting.AppSetting.LogSaveName,
		setting.AppSetting.LogFileExt,
		time.Now().Format(setting.AppSetting.TimeFormat),
	)
}

func logFileDir(rootPath string, dirName string) string {
	return fmt.Sprintf("%s%s", rootPath, dirName)
}

func logFileName(file string, ext string, date string) string {
	return fmt.Sprintf("%s.%s.%s",
		file,
		ext,
		date,
	)
}
