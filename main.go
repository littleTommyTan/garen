package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/raven-go"
	"github.com/tommytan/garen/internal/justice"
	"github.com/tommytan/garen/internal/service"

	"github.com/tommytan/garen/configs"
	goCron "github.com/tommytan/garen/internal/cron"
	"github.com/tommytan/garen/internal/judgment"
)

var appYaml = flag.String("configuration", "configs/conf.yaml", "garen judgment configuration file")

// init 系统函数 初始化
func init() {
	flag.Parse()
	if err := configs.LoadConfiguration(*appYaml); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	_ = raven.SetDSN("https://06590039334549c18b926a9623b7c183:d64f1d94d5b94a2a8915b3c42f188980@sentry.io/1461854")
}

// main 入口 entrance
func main() {
	// 阿里云日志服务测试
	//logger.AliyunLoggerTest()

	// 初始化定时任务
	goCron.Cron()

	// 初始化系统服务
	service.New()

	// 正义 grpc server
	jGRPC := justice.SetupGrpcJustice()

	// 审判 http server
	jHTTP := judgment.SetupHttpJudgment()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			jGRPC.Stop()
			_ = jHTTP.Close()
			service.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
