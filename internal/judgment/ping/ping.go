package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/judgment/jwt"
	"log"
	"net/http"
	"time"
)

func Assemble(r *gin.Engine) {
	g := r.Group("/ping")
	{
		g.GET("/hello", func(c *gin.Context) {
			c.String(200, "pong")
		})

		g.GET("/hello/:hi/nihao/:lalala", func(c *gin.Context) {
			hi := c.Param("hi")
			lalala := c.Param("lalala")
			if lalala == "" || hi == "" {
				c.String(400, "params validation failed. ")
				return
			}
			c.String(200, `hi, params ping. hi:%v lalala:%v`, hi, lalala)
		})

		g.GET("/vip", jwt.Auth(), func(c *gin.Context) {
			val, _ := c.MustGet("userId").(string)
			c.String(http.StatusOK, "pong pong pong, hi vip "+val)
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
