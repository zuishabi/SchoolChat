package routers

import (
	"SchoolChat/database"
	"SchoolChat/globals"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := c.Query("uid")
	if uid == "" {
		c.Redirect(http.StatusSeeOther, "/userInfo?uid="+session.Get("uid").(string))
		return
	}
	c.HTML(http.StatusOK, "user_info.html", nil)
}

type UserInfoReq struct {
	Name         string `json:"name"`
	CurrentUID   string `json:"current_uid"`
	ID           string `json:"student_id"`
	GraduateTime string `json:"graduate_time"`
	Description  string `json:"description"`
}

func GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	sourceUID := session.Get("uid").(string)
	uid := c.Query("uid")
	if uid == "" {
		uid = sourceUID
	}
	user := database.UserInfo{}
	database.Db.Where("uid = ?", uid).First(&user)
	res := UserInfoReq{
		Name:         user.UserName,
		CurrentUID:   sourceUID,
		ID:           user.UID,
		GraduateTime: user.GraduateTime,
		Description:  user.Description,
	}
	c.JSON(http.StatusOK, &res)
}

// UpdateUserInfo 更新用户的信息
func UpdateUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	info := globals.UpdateUserInfoReq{}
	if err := c.BindJSON(&info); err != nil {
		fmt.Println(err)
		c.String(http.StatusForbidden, "解析json错误")
		return
	}
	database.Db.Model(database.UserInfo{}).Where("uid = ?", session.Get("uid").(string)).Updates(database.UserInfo{
		UserName:     info.Name,
		GraduateTime: info.GraduateTime,
		Description:  info.Description,
	})
	globals.AddUpdate(session.Get("uid").(string), info)
	session.Set("user_name", info.Name)
	_ = session.Save()
	c.String(http.StatusOK, "修改成功")
}
