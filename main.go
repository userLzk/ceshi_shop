package main

import (
	_ "net/http/pprof"

	"ceshi_shop/route"
	"github.com/gin-gonic/gin"
)

//创建链接初始化信息
func init() {
	//创建链接
	//Connection.ConnectionStart()
}

func main() {
	//创建路由引擎
	r := gin.Default()
	route.IninServer(r)
	gin.SetMode(gin.ReleaseMode)
	r.Run(":9999")

}
