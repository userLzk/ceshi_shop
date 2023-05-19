package route

import (
	"ceshi_shop/server"
	"github.com/gin-gonic/gin"
)

func IninServer(r *gin.Engine) {
	//公共检测

	//路由注册
	login := r.Group("/server")
	{
		login.POST("/login", server.Login)
		login.POST("/getToken", server.GetToken)
	}
}
