package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/tommytan/garen/internal/service"
)

// AuthRequired is jwt auth middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-garen-token")
		if token == "" {
			c.String(401, "Not Authorized")
			c.Abort()
			return
		}

		// 实际上，Dao结构体的redis和db对象不应该被导出，所以应该在Dao里面写关于数据访问的逻辑就像下面的service.Dao.VerifyToken，这是因为数据访问的逻辑大多可以复用，如getList
		// 不过，因为我们项目小且需要快速开发，允许一些较简单的访问逻辑，直接处理就好

		//openId, err := service.Dao.Redis.Get("userlist:" + token).Result()
		openId, err := service.Dao.VerifyToken("userlist:" + token)

		if err == redis.Nil {
			c.String(401, "Not Authorized")
			c.Abort()
			return
		}

		// TODO 刷新

		c.Set("userId", openId)
		c.Next()
	}
}
