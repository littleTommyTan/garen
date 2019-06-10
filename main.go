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
	"github.com/littletommytan/garen/internal/justice"
	"github.com/littletommytan/garen/internal/service"

	"github.com/littletommytan/garen/configs"
	goCron "github.com/littletommytan/garen/internal/cron"
	"github.com/littletommytan/garen/internal/judgment"
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
	jHTTP := judgment.SetupHTTPJudgment()

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
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

//package main
//
//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/littletommytan/garen/internal/gokit/greeterendpoint"
//	"github.com/littletommytan/garen/internal/gokit/greetersd"
//	"github.com/littletommytan/garen/internal/gokit/greeterservice"
//	"github.com/littletommytan/garen/internal/gokit/greetertransport"
//	"github.com/littletommytan/garen/internal/judgment/logger"
//	"github.com/sirupsen/logrus"
//	"os/signal"
//	"syscall"
//	"time"
//
//	"google.golang.org/grpc"
//	"net"
//	"net/http"
//	"os"
//)
//
//func main() {
//
//	r := gin.Default()
//	r.Use(logger.LocalFileLogger())
//	r.GET("/health", func(context *gin.Context) {
//		context.String(200, "healthy!")
//	})
//	r.GET("/greeting", func(context *gin.Context) {
//		context.String(200, "greeting!")
//	})
//	httpserver := &http.Server{
//		Addr:           ":5000",
//		Handler:        r,
//		ReadTimeout:    5 * time.Second,
//		WriteTimeout:   5 * time.Second,
//		MaxHeaderBytes: 2 << 20,
//	}
//
//	service := greeterservice.GreeterService{}
//
//	var (
//		endpoints  = greeterendpoint.MakeServerEndpoints(service)
//		registar   = greetersd.ConsulRegister()
//		grpcServer = greetertransport.NewGRPCServer(endpoints)
//	)
//
//	go func() {
//		registar.Register()
//		logrus.Infof("consul on http://localhost:8500")
//		logrus.Infof("http server is running on http://localhost:5000")
//		if err := httpserver.ListenAndServe(); err != nil {
//			registar.Deregister()
//			logrus.Error(err)
//			os.Exit(1)
//		}
//	}()
//
//	var baseServer *grpc.Server
//	go func() {
//		grpcListener, err := net.Listen("tcp", ":9120")
//		if err != nil {
//			logrus.Error(err)
//			os.Exit(1)
//		}
//		baseServer = grpc.NewServer()
//		logrus.Infof("grpc server is running ... ")
//		greetertransport.RegisterGreeterServer(baseServer, grpcServer)
//		_ = baseServer.Serve(grpcListener)
//	}()
//
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
//	for {
//		s := <-c
//		switch s {
//		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
//			_, cancel := context.WithTimeout(context.Background(), 35*time.Second)
//			baseServer.Stop()
//			_ = httpserver.Close()
//			registar.Deregister()
//			cancel()
//			time.Sleep(time.Second)
//			return
//		case syscall.SIGHUP:
//		default:
//			return
//		}
//	}
//}
