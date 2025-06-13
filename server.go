package main

import (
	"SchoolChat/globals"
	"SchoolChat/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	globals.InitBloomFilter()
	store := cookie.NewStore([]byte("secret"))
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", routers.Login)
	r.POST("/", routers.LoginHandler)
	r.GET("/register", routers.Register)
	r.POST("register", routers.RegisterHandler)
	r.GET("/main", routers.Main)
	r.POST("/sendChat", routers.SendChatHandler)
	r.GET("/chatBoard", routers.ChatBoard)
	r.GET("/logout", routers.Logout)
	r.GET("/createClass", routers.CreateClass)
	r.GET("/getClass", routers.GetClass)
	r.GET("/getClassMembers", routers.GetClassMembers)
	r.Run(":8888")
}
