package routers

import (
	"SchoolChat/database"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UID      string `json:"uid"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	s := sessions.Default(c)
	if s.Get("uid") != nil {
		fmt.Println("get session ", s.Get("uid").(string))
		c.Redirect(http.StatusSeeOther, "/main")
		return
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginHandler(c *gin.Context) {
	loginInfo := LoginInfo{}
	if err := c.BindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"err": "解析登录信息错误"})
		return
	}
	userInfo := database.UserInfo{}
	err := database.Db.Where("uid = ? and password = ?", loginInfo.UID, loginInfo.Password).First(&userInfo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "账号或密码错误"})
		return
	}
	session := sessions.Default(c)

	// 设置session值
	session.Set("uid", userInfo.UID)
	session.Set("user_name", userInfo.UserName)
	session.Set("graduate_time", userInfo.GraduateTime)
	// 设置session过期时间(可选)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24, // 24小时
	})

	// 保存session
	err = session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/main")
}
