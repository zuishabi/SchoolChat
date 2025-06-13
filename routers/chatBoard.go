package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	fmt.Println(page, uid)
}
