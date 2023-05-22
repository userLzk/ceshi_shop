package route

import (
	"ceshi_shop/server"
	"ceshi_shop/verfiy"
	"github.com/gin-gonic/gin"
)

func IninServer(r *gin.Engine) {

	//路由注册
	login := r.Group("/server")
	{
		login.POST("/login", server.Login)
		login.POST("/getToken", verfiy.JwtAuth(), server.GetToken)
	}
}
