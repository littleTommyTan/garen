package justice

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
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
