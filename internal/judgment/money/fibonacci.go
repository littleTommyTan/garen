package money

import "github.com/gin-gonic/gin"

func calFibonacci(c *gin.Context) {
	ch := make(chan int, 15) //定义channel缓冲大小时为15，表示最多存取15个元素。
	go func() {
		x, y := 0, 1
		for i := 0; i < 15; i++ {
			ch <- x
			x, y = y, x+y
			if x > 100 { //表示当循环到两个数字之和大于100时就终止循环。
				break
			}
		}
		close(ch) //关闭channel,如果不关闭就会被锁，出现报错："deadlock!"
	}()
	sum := 0
	for i := range ch { //遍历channel_name
		sum += i
	}
	c.String(200, "fibonacci cal done %v", sum)
}
