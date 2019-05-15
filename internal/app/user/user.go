package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tommytan/garen/internal/middleware/jwt"
	"github.com/tommytan/garen/internal/models"
	"github.com/tommytan/garen/internal/service"
	"log"
	"net/http"
	"time"
)

func DecorateRouterGroup(r *gin.Engine) {
	g := r.Group("/user")
	{
		g.GET("/register", register)
		g.GET("/login", login)
	}
}

func register(c *gin.Context) {
	var form struct {
		Nickname string `form:"nickname" binding:"required"`
	}
	err := c.ShouldBindQuery(&form)
	if err != nil {
		c.String(400, "参数不正确")
		return
	}

	service.Dao.Db.Create(&models.User{NickName: form.Nickname})
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

	token, _ := jwt.NewJWT().CreateToken(form.Nickname)
	service.Dao.Redis.Set("userlist:"+token, form.Nickname, time.Hour*2)
	c.String(200, token)
}
