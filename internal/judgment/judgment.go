package judgment

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/judgment/middleware/logger"
	"github.com/tommytan/garen/internal/judgment/music"
	"github.com/tommytan/garen/internal/judgment/ping"
	"github.com/tommytan/garen/internal/judgment/user"
	"github.com/tommytan/garen/internal/judgment/wechat"
	"github.com/tommytan/garen/internal/judgment/whoareyou"
	"log"
	"net/http"
	"time"
)

// SetupHttpJudgment 路由设置router
func SetupHttpJudgment() (s *http.Server) {
	r := gin.New()

	r.Use(sentry.Recovery(raven.DefaultClient, false))

	r.Use(logger.LocalFileLogger())

	ping.Assemble(r)

	wechat.Assemble(r)

	user.Assemble(r)

	music.Assemble(r)

	whoareyou.Assemble(r)

	s = &http.Server{
		Addr:           ":2333",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	return
}
