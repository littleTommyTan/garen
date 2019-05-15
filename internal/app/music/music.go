package music

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func DecorateRouterGroup(r *gin.Engine) {
	g := r.Group("/music")
	{
		g.GET("/hello", GetSongName)
	}
}

// GetSongName 获取你的歌单里的第一首歌的名字 （interface map 强制类型转换入门）
//
// test: curl "http://localhost:2333/music/hello?listId=169174938"
func GetSongName(c *gin.Context) {
	listId := c.Query("listId")
	if listId == "" {
		c.String(400, "参数校验错误")
	}

	resp, err := http.Get(fmt.Sprint("https://bird.ioliu.cn/netease/playlist?id=", listId))
	if err != nil {
		c.String(http.StatusServiceUnavailable, "")
	}

	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("json marshal failed,err:", err)
		return
	}
	playlist := data["playlist"].(map[string]interface{})
	tracks := playlist["tracks"].([]interface{})
	song := tracks[0].(map[string]interface{})
	songName := song["name"].(string)

	c.String(http.StatusOK, songName)
}
