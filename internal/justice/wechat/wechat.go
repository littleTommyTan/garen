package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/tommytan/garen/configs"
	"net/http"
)

var wechatService *wechat.Wechat

func Assemble(r *gin.Engine) {
	redis := cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379", Password: "redis-pwd", Database: 0, MaxIdle: 0, MaxActive: 0, IdleTimeout: 0})
	//配置微信参数
	config := &wechat.Config{
		AppID:     configs.GetConfiguration().WoaAppId,
		AppSecret: configs.GetConfiguration().WoaAppSecret,
		//Token:          "your token",
		//EncodingAESKey: "your encoding aes key",
		Cache: redis,
	}
	wechatService = wechat.NewWechat(config)

	g := r.Group("/wechat")
	{
		g.GET("/conf", getJsConf)
	}
}

func getJsConf(c *gin.Context) {
	config, err := wechatService.GetJs().GetConfig(c.Request.URL.Path)
	if err != nil {
		c.String(http.StatusServiceUnavailable, "get js conf failed")
		return
	}
	c.SecureJSON(200, gin.H{
		"debug":     false,
		"appId":     config.AppID,
		"timestamp": config.Timestamp,
		"nonceStr":  config.NonceStr,
		"signature": config.Signature,
		"jsApiList": []string{"getNetworkType", "updateAppMessageShareData", "updateTimelineShareData"},
	})
}
