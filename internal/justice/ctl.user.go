package justice

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/models"
	"log"
	"net/http"
	"time"
)

func register(c *gin.Context) {
	var form struct {
		Nickname string `form:"nickname" binding:"required"`
	}
	err := c.ShouldBindQuery(&form)
	if err != nil {
		c.String(400, "参数不正确")
		return
	}

	svc.Dao.Db.Create(&models.User{NickName: form.Nickname})
	c.String(http.StatusOK, form.Nickname+" 添加成功")
}

func login(c *gin.Context) {
	var form struct {
		Nickname string `form:"nickname" binding:"required"`
	}
	err := c.ShouldBindQuery(&form)
	if err != nil {
		log.Print(err.Error())
		c.String(400, "参数不正确")
		return
	}

	token, _ := NewJWT().CreateToken(form.Nickname)
	svc.Dao.Redis.Set("userlist:"+token, form.Nickname, time.Hour*2)
	c.String(200, token)
}
