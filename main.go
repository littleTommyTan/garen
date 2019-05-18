package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/raven-go"
	"net/http"
	"os"
	"time"

	"github.com/tommytan/garen/configs"
	goCron "github.com/tommytan/garen/internal/cron"
	"github.com/tommytan/garen/internal/justice"
	"github.com/tommytan/garen/internal/service"
)

var appYaml = flag.String("configuration", "configs/conf.yaml", "garen justice configuration file")

// init 系统函数 初始化
func init() {
	flag.Parse()
	if err := configs.LoadConfiguration(*appYaml); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

// main 入口 entrance
func main() {

	// Sentry
	_ = raven.SetDSN("https://06590039334549c18b926a9623b7c183:d64f1d94d5b94a2a8915b3c42f188980@sentry.io/1461854")

	// 初始化定时任务
	goCron.Cron()

	// 正义 setupRouter
	j := justice.SetupJustice(service.New())

	s := &http.Server{
		Addr:           ":2333",
		Handler:        j,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	_ = s.ListenAndServe()
}
