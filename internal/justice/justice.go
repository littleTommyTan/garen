package justice

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/justice/music"
	"github.com/tommytan/garen/internal/justice/ping"
	"github.com/tommytan/garen/internal/justice/user"
	"github.com/tommytan/garen/internal/justice/wechat"
	"github.com/tommytan/garen/internal/justice/whoareyou"
)

// SetupJudgment 路由设置router
func SetupJudgment() *gin.Engine {
	r := gin.Default()

	ping.Assemble(r)

	wechat.Assemble(r)

	user.Assemble(r)

	music.Assemble(r)

	whoareyou.Assemble(r)

	return r
}
