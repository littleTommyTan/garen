package judgment

import (
	"log"
	"net/http"
	"time"

	"github.com/tommytan/garen/internal/judgment/goroutine"

	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tommytan/garen/internal/judgment/logger"
	"github.com/tommytan/garen/internal/judgment/money"
	"github.com/tommytan/garen/internal/judgment/music"
	"github.com/tommytan/garen/internal/judgment/ping"
	"github.com/tommytan/garen/internal/judgment/user"
	"github.com/tommytan/garen/internal/judgment/wechat"
	"github.com/tommytan/garen/internal/judgment/whoareyou"
)

// SetupHTTPJudgment 路由设置router
func SetupHTTPJudgment() (s *http.Server) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Use(logger.LocalFileLogger())
	//r.Use(iplimiter.NewRateLimiterMiddleware(redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%v:%v", configs.GetConfiguration().RedisHost, configs.GetConfiguration().RedisPort),
	//	Password: configs.GetConfiguration().RedisPwd,
	//	DB:       1,
	//}), "tommy-iplimiter", 100, 60*time.Second))

	ping.Assemble(r)

	wechat.Assemble(r)

	user.Assemble(r)

	music.Assemble(r)

	whoareyou.Assemble(r)

	money.Assemble(r)

	goroutine.Assemble(r)

	s = &http.Server{
		Addr:           ":2333",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	logrus.Infof("http server is running ...")
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	return
}
