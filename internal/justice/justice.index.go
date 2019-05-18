package justice

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/service"
)

var (
	svc *service.Service
)

// SetupJustice 路由设置router
func SetupJustice(s *service.Service) *gin.Engine {
	svc = s
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.Use(JwtAuth())

	pingGroup := r.Group("/ping")
	{
		pingGroup.GET("/hello", normalPing)
		pingGroup.GET("/vip", JwtAuth(), middlewarePing)
		pingGroup.GET("/long_async", goRoutinePing)
	}

	musicGroup := r.Group("/music")
	{
		musicGroup.GET("/hello", GetSongName)
	}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/register", register)
		userGroup.GET("/login", login)
	}

	wechatGroup := r.Group("/wechat")
	{
		wechatGroup.GET("/conf", getJsConf)
	}

	whoareyouGroup := r.Group("/whoareyou")
	{
		whoareyouGroup.GET("/hello", hello)
		whoareyouGroup.POST("/uploadAvatar", uploadAvatar)
		whoareyouGroup.POST("/uploadAvatarSmms", uploadAvatarSmms)
	}

	return r
}
