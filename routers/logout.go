package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		fmt.Println("save session err = ", err)
	}
	c.Redirect(http.StatusSeeOther, "/")
}
