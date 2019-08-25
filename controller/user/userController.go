package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.tmpl", gin.H{
		"title": "user",
	})
}
