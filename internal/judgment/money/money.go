package money

import (
	"github.com/gin-gonic/gin"
)

// Assemble 路由集结
func Assemble(r *gin.Engine) {
	g := r.Group("/money")
	{
		g.GET("/hello", hello)
		g.GET("/mutex", dosomethingWithMyMoney)
		g.GET("/fi", calFibonacci)
		g.GET("/select", cal)
		g.GET("/default", defaultss)
	}
}
