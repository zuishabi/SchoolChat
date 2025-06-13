package routers

import (
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
