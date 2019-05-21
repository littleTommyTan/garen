package justice

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/justice/middleware/logger"
	"github.com/tommytan/garen/internal/justice/music"
	"github.com/tommytan/garen/internal/justice/ping"
	"github.com/tommytan/garen/internal/justice/user"
	"github.com/tommytan/garen/internal/justice/wechat"
	"github.com/tommytan/garen/internal/justice/whoareyou"
)

// SetupJudgment 路由设置router
func SetupJudgment() *gin.Engine {
	r := gin.New()

	r.Use(sentry.Recovery(raven.DefaultClient, false))

	r.Use(logger.LocalFileLogger())

	ping.Assemble(r)

	wechat.Assemble(r)

	user.Assemble(r)

	music.Assemble(r)

	whoareyou.Assemble(r)

	return r
}
