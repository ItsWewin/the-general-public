package router

import (
	"github.com/ItsWewin/the-general-public/controller/session"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("/Users/klook/workspace/src/github.com/ItsWewin/the-general-public/view/*/*")
	router.GET("/login", session.Login)

	return router
}
