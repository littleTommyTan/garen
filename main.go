package main

import (
	"flag"
	"fmt"
	"github.com/tommytan/garen/configs"
	"github.com/tommytan/garen/internal/justice"
	"github.com/tommytan/garen/internal/service"
	"net/http"
	"os"
	"time"
)

var appYaml = flag.String("configuration", "configs/conf.yaml", "garen justice configuration file")

// main 入口 entrance
func main() {
	flag.Parse()

	if err := configs.LoadConfiguration(*appYaml); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//gin.SetMode(gin.ReleaseMode)

	service.New()

	//初始化路由 setupRouter
	j := justice.SetupJustice()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        j,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
