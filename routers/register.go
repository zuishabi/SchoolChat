package routers

import (
	"SchoolChat/database"
	"SchoolChat/globals"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RegisterInfo struct {
	UID          string `json:"uid"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	GraduateTime string `json:"graduate_time"`
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterHandler(c *gin.Context) {
	registerInfo := RegisterInfo{}
	if err := c.BindJSON(&registerInfo); err != nil {
		fmt.Println(err)
		return
	}
	// 查询学号是否重复
	userInfo := database.UserInfo{
		UID:          registerInfo.UID,
		UserName:     registerInfo.UserName,
		Password:     registerInfo.Password,
		GraduateTime: registerInfo.GraduateTime,
	}
	if !globals.B.GetItem(registerInfo.UID) {
		database.Db.Create(&userInfo)
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	u := database.UserInfo{}
	if err := database.Db.Where("uid = ?", userInfo.UID).First(&u).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		database.Db.Create(&userInfo)
		c.JSON(http.StatusOK, gin.H{})
		return
	} else {
		fmt.Println("查询用户信息错误,err = ", err)
		c.JSON(http.StatusForbidden, gin.H{"err": "当前学号已被注册"})
		return
	}
}
