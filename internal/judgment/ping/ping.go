package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/judgment/middleware/jwt"
	"log"
	"time"
)

func Assemble(r *gin.Engine) {
	g := r.Group("/ping")
	{
		g.GET("/hello", func(c *gin.Context) {
			c.String(200, "pong")
		})

		g.GET("/vip", jwt.Auth(), func(c *gin.Context) {
			val, _ := c.MustGet("userId").(string)
			c.String(200, "pong pong pong, hi vip "+val)
		})

		g.GET("/long_async", func(c *gin.Context) {
			// 创建要在goroutine中使用的副本
			cCp := c.Copy()
			go func() {
				// simulate a long task with time.Sleep(). 5 seconds
				time.Sleep(2 * time.Second)

				// 这里使用你创建的副本
				log.Println("Done! in path " + cCp.Request.URL.Path)
			}()
		})
	}

}
