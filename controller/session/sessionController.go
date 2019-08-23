package session

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "session/login.tmpl", gin.H{
		"title": "users",
	})
}