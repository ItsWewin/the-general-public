package router

import (
	"github.com/ItsWewin/the-general-public/controller/session"
	"github.com/ItsWewin/the-general-public/controller/user"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("/Users/wewin/goprojects/src/github.com/ItsWewin/the-general-public/view/*/*")
	router.GET("/login", session.Login)

	router.GET("/join", user.Register)

	return router
}
