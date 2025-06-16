package routers

import (
	"SchoolChat/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Chat struct {
	Messages []Message `json:"messages"`
	Total    int64     `json:"total_pages"`
}

type Message struct {
	Name    string `json:"from_name"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

func ChatBoard(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	uid := c.Query("uid")
	// 如果是第一页就先在缓存中查询
	// 否则在数据库中查询
	chats := make([]database.ChatBoard, 0)
	var total int64
	database.Db.Model(database.ChatBoard{}).Where("uid = ?", uid).Count(&total)
	database.Db.Where("uid = ?", uid).Order("uid desc").Offset((page - 1) * 10).Limit(10).Find(&chats)
	messages := make([]Message, len(chats))
	for i, v := range chats {
		messages[i].Name = v.SenderName
		messages[i].Time = v.CreatedAt.String()
		messages[i].Content = v.Content
	}
	if len(chats) == 0 {
		c.String(http.StatusForbidden, "没有对应数据")
		return
	}
	c.JSON(http.StatusOK, Chat{
		Messages: messages,
		Total:    total/11 + 1,
	})
}

type sendChat struct {
	UID     string `json:"uid"`
	Content string `json:"content"`
}

// CreateChat 创建一个留言
func CreateChat(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	// 发送者的uid
	name := session.Get("user_name").(string)
	uid := session.Get("uid").(string)
	send := sendChat{}
	if err := c.BindJSON(&send); err != nil {
		c.String(http.StatusForbidden, err.Error())
		return
	}
	database.Db.Create(&database.ChatBoard{
		UID:        send.UID,
		SenderName: name,
		Content:    send.Content,
		Sender:     uid,
	})
	c.String(http.StatusOK, "发送成功")
}
