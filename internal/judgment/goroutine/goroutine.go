package goroutine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Assemble 路由集结
func Assemble(r *gin.Engine) {
	g := r.Group("/goroutine")
	{
		g.GET("/hello", hello)
	}
}
func sayhi(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func hello(c *gin.Context) {
	s := c.Query("hi")
	if s == "" {
		c.String(400, "params error")
		return
	}
	c.String(200, s)
	go sayhi("Tommy Tan")
	sayhi("Golang")
}
