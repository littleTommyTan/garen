package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logClient = logrus.New()

func LocalFileLogger() gin.HandlerFunc {
	//src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	_, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//logClient.Out = src
	logClient.Out = os.Stdout

	//logClient.SetLevel(logrus.DebugLevel)

	//log writer
	logPath := "garen.log"
	logWriter, err := rotatelogs.New(
		"garen.%Y-%m-%d.log",
		rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	//local file system hook
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		// 其他参数
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logClient.WithFields(logrus.Fields{"nihao": "hello", "nishishei": "whoareyou"}).Debugf("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
		logClient.WithFields(logrus.Fields{"nihao": "hello", "nishishei": "whoareyou"}).Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
		logClient.WithFields(logrus.Fields{"nihao": "hello", "nishishei": "whoareyou"}).Warnf("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path)
		logClient.WithFields(logrus.Fields{"nihao": "hello", "nishishei": "whoareyou"}).Errorf("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path)
	}
}
