package justice

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func normalPing(c *gin.Context) {
	c.String(200, "pong")
}

func middlewarePing(c *gin.Context) {
	val, _ := c.MustGet("userId").(string)
	c.String(200, "pong pong pong, hi vip "+val)
}

func goRoutinePing(c *gin.Context) {
	// 创建要在goroutine中使用的副本
	cCp := c.Copy()
	go func() {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(2 * time.Second)

		// 这里使用你创建的副本
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
