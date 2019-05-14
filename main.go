package main

import (
	"flag"
	"fmt"
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
	// 初始化系统服务
	service.New()
}

// main 入口 entrance
func main() {

	//gin.SetMode(gin.ReleaseMode)

	// 初始化定时任务
	goCron.Cron()

	// 正义 setupRouter
	j := justice.SetupJustice()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        j,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	_ = s.ListenAndServe()
}
