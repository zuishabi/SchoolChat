package routers

import (
	"SchoolChat/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type chatInfo struct {
	UID     string `json:"uid"`
	Content string `json:"content"`
}

func SendChatHandler(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uid := session.Get("uid").(string)
	chat := chatInfo{}
	err := c.BindJSON(&chat)
	if err != nil {
		c.String(http.StatusForbidden, "解析json失败")
		return
	}
	// 将数据存储到数据库中
	b := database.ChatBoard{
		UID:     chat.UID,
		Sender:  uid,
		Content: chat.Content,
	}
	database.Db.Create(&b)
	// 将数据库中的数据删除
}
