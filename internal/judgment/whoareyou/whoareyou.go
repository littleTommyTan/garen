package whoareyou

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/service"
	"net/http"
	"time"
)

func Assemble(r *gin.Engine) {
	g := r.Group("/whoareyou")
	{
		g.GET("/hello", hello)
		g.GET("/grpc-hello", grpcHello)
		g.POST("/uploadAvatar", uploadAvatar)
		g.POST("/uploadAvatarSmms", uploadAvatarSmms)
	}
}

func uploadAvatar(c *gin.Context) {
	// 单文件
	file, fh, err := c.Request.FormFile("file")
	if err != nil {
		c.String(400, "file invalid")
		return
	}
	url, err := service.BucketUpload(file, fh)
	if err != nil {
		c.String(http.StatusServiceUnavailable, fmt.Sprintf("'%s' upload failed!", fh.Filename))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", url))
}

func uploadAvatarSmms(c *gin.Context) {
	// 单文件
	file, fh, err := c.Request.FormFile("file")
	if err != nil {
		c.String(400, "file invalid")
		return
	}
	url, err := service.SmmsUpload(file, fh)
	if err != nil {
		c.String(http.StatusServiceUnavailable, fmt.Sprintf("'%s' upload failed!", fh.Filename))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", url))
}

func hello(c *gin.Context) {
	birthday := c.Query("birthday")
	y, m, d, err := getTimeFromStrDate(birthday)
	if err != nil {
		c.String(http.StatusBadRequest, "出生日期格式不正确")
		raven.CaptureError(err, map[string]string{"category": "test"})
		return
	}
	c.SecureJSON(200, gin.H{
		"age":           getAge(y),
		"constellation": getConstellation(m, d),
		"zodiac":        getZodiac(y),
	})
}

func getZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func getAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

func getConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}

	return star
}

func getTimeFromStrDate(date string) (int, int, int, error) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		return 0, 0, 0, err
	}
	return d.Year(), int(d.Month()), d.Day(), err
}
