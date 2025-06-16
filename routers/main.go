package routers

import (
	"SchoolChat/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	s := sessions.Default(c)
	if s.Get("uid") == nil {
		c.Redirect(http.StatusUseProxy, "/")
		return
	}
	c.HTML(http.StatusOK, "main.html", nil)
}

type SearchInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	GraduateTime string `json:"graduate_time"`
}

// SearchUser 通过查询信息检索对应用户
func SearchUser(c *gin.Context) {
	s := sessions.Default(c)
	if s.Get("uid") == nil {
		c.Redirect(http.StatusUseProxy, "/")
		return
	}
	search := database.UserInfo{
		UID:          c.Query("uid"),
		UserName:     c.Query("user_name"),
		GraduateTime: c.Query("graduate_time"),
	}
	res := make([]database.UserInfo, 0)
	database.Db.Where(&search).Find(&res)
	infos := make([]SearchInfo, len(res))
	for i, v := range res {
		infos[i].Name = v.UserName
		infos[i].ID = v.UID
		infos[i].Description = v.Description
		infos[i].GraduateTime = v.GraduateTime
	}
	c.JSON(http.StatusOK, &infos)
}
