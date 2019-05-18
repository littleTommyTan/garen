package justice

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/justice/music"
	"github.com/tommytan/garen/internal/justice/ping"
	"github.com/tommytan/garen/internal/justice/user"
	"github.com/tommytan/garen/internal/justice/wechat"
	"github.com/tommytan/garen/internal/justice/whoareyou"
)

// SetupJustice 路由设置router
func SetupJustice() *gin.Engine {
	r := gin.Default()

	ping.DecorateRouterGroup(r)

	wechat.DecorateRouterGroup(r)

	user.DecorateRouterGroup(r)

	music.DecorateRouterGroup(r)

	whoareyou.DecorateRouterGroup(r)

	return r
}
