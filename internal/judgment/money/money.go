package money

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Assemble 路由集结
func Assemble(r *gin.Engine) {
	g := r.Group("/money")
	{
		g.GET("/hello", hello)
		g.GET("/do", dosomethingWithMyMoney)
	}
}

type Request struct {
	data []int
	ret  chan int
}

func NewRequest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}

func hello(c *gin.Context) {
	req := NewRequest(10, 20, 30)
	Process(req)
	c.String(http.StatusOK, "hello money %v", <-req.ret)
}

var m *sync.RWMutex //定义m的类型为读写锁。

func read(i string) {
	println(i, "read start")
	m.RLock()
	println(i, "资料正在读取中.....")
	time.Sleep(4 * time.Second)
	m.RUnlock()
	println(i, "read end")
}

func write(i string) {
	println(i, "write start")
	m.Lock() //写操作锁定
	println(i, "数据正在写入硬盘中.....")
	time.Sleep(4 * time.Second)
	m.Unlock() //写操作解锁
	println(i, "write end")
}

func dosomethingWithMyMoney(c *gin.Context) {
	m = new(sync.RWMutex)
	go write("第一次写入") //写的时候啥都不能干,即其他协程无法写入，也无法读取。
	go read("第一次读取")
	go write("第二次写入")
	time.Sleep(15 * time.Second) //主进程只给出10秒的时间，但是整个进程跑完最少需要12秒的时间。
	c.String(200, "")
}
