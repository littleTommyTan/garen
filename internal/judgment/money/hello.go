package money

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	data []int
	ret  chan int
}

func newRequest(data ...int) *request {
	return &request{data, make(chan int, 1)}
}

func process(req *request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}

func hello(c *gin.Context) {
	req := newRequest(10, 20, 30)
	process(req)
	c.String(http.StatusOK, "hello money %v", <-req.ret)
}
