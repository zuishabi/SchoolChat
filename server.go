package main

import (
	"SchoolChat/globals"
	"SchoolChat/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	globals.InitUpdater()
	globals.InitBloomFilter()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "", // 或留空
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
	})
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", routers.Login)
	r.POST("/", routers.LoginHandler)
	r.GET("/register", routers.Register)
	r.POST("register", routers.RegisterHandler)
	r.GET("/main", routers.Main)
	r.GET("/logout", routers.Logout)
	r.GET("/createClass", routers.CreateClass)
	r.GET("/getClass", routers.GetClass)
	r.GET("/getClassMembers", routers.GetClassMembers)
	r.GET("/manageClass", routers.ManageClass)
	r.GET("/searchClass", routers.SearchClass)
	r.GET("/addClass", routers.AddClass)
	r.GET("/deleteClass", routers.DeleteClass)
	r.GET("/getEnterRequest", routers.GetEnterRequest)
	r.GET("/approveEnter", routers.ApproveEnter)
	r.GET("/rejectEnter", routers.RejectEnter)
	r.GET("/kickMember", routers.RemoveUser)
	r.GET("/quitClass", routers.QuitClass)
	r.GET("/userInfo", routers.UserInfo)
	r.GET("/getUserInfo", routers.GetUserInfo)
	r.POST("/updateUserInfo", routers.UpdateUserInfo)
	r.GET("/getMessages", routers.ChatBoard)
	r.POST("/leaveMessage", routers.CreateChat)
	r.GET("/searchUser", routers.SearchUser)
	r.Run(":7878")
}
