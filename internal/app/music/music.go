package music

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func DecorateRouterGroup(r *gin.Engine) {
	g := r.Group("/music")
	{
		g.GET("/hello", FetchMusicList)
	}
}

func FetchMusicList(c *gin.Context) {
	resp, err := http.Get("https://bird.ioliu.cn/netease/playlist?id=169174938")
	if err != nil {
		c.String(http.StatusServiceUnavailable,"")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusServiceUnavailable,"")
	}
	c.String(http.StatusOK,string(body))
}
