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
)

// SetupHttpJudgment 路由设置router
func SetupHttpJudgment() *gin.Engine {
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
