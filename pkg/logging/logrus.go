package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	//now := time.Now()
	//logFilePath := getLogFilePath()
	//logFileName :=  getLogFileFullPath()
	//日志文件
	//fileName := path.Join(logFilePath, logFileName)
	//if _, err := os.Stat(fileName); err != nil {
	//	if _, err := os.Create(fileName); err != nil {
	//		fmt.Println(err.Error())
	//	}
	//}
	//写入文件
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err := openLogFile(fileName, filePath)
	if err != nil {
		fmt.Println("日志写入失败，请检查")
	}
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = F

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
