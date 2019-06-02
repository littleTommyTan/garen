package money

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func fibonacci(ch, quit chan int) { //定义两个channle对象channel_name和quit。
	x, y := 0, 1
	for {
		select {
		case ch <- x: //用channel_name接受数据。
			x, y = y, x+y

		case <-quit: //表示当接收到quit的channel时，就执行以下代码。其实就是实现关闭channel的功能。
			fmt.Println("EXIT")
			return //函数一退出协程也就跟着退出了
		}
	}
}

func cal(c *gin.Context) {
	tick := time.Tick(1000 * time.Millisecond) //也可以这样写：“tick := time.NewTicker(1000*time.Millisecond).C”其中这个点C就是一个channel。
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("dida.")
		case <-boom:
			fmt.Println("boom.")
			c.String(200, "done %v", "test")
			return
		default:
			fmt.Println("say hi.")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// ch := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 11; i++ {
	// 		fmt.Println(<-ch)
	// 	}
	// 	quit <- 100

	// }()
	// time.Sleep(time.Second * 3)
	// fibonacci(ch, quit)
}
