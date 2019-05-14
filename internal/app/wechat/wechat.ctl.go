package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"log"
	"net/http"
)

var wechatService *wechat.Wechat

func init() {
	redis := cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379", Password: "redis-pwd", Database: 0, MaxIdle: 0, MaxActive: 0, IdleTimeout: 0})
	//配置微信参数
	config := &wechat.Config{
		AppID:     "wx6a042d99b5df1a93",
		AppSecret: "f0eed325aa629a74581b4b66591e85f0",
		//Token:          "your token",
		//EncodingAESKey: "your encoding aes key",
		Cache: redis,
	}
	wechatService = wechat.NewWechat(config)
}

func DecorateRouterGroup(r *gin.Engine) {
	g := r.Group("/wechat")
	{
		g.GET("/hello", hello)
	}
}

func hello(c *gin.Context) {
	accessToken, err := wechatService.GetAccessToken()
	if err != nil {
		log.Print(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	c.String(http.StatusOK, accessToken)
}

func FetchAccessToken() (string, error) {
	accessToken, err := wechatService.GetAccessToken()
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	log.Print(accessToken)
	return accessToken, nil
}
