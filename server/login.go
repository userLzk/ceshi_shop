package server

import (
	"fmt"

	"ceshi_shop/pkg/jwt"
	"github.com/gin-gonic/gin"
)

//创建对象

var LoginController loginController

type loginController struct {
}

func Login(g *gin.Context) {
	car := jwt.NewCarryData("zhagnsan", 17)
	token := jwt.EncryptionToken(car)
	tokenDesc, _ := jwt.DecryptionToken(token)
	fmt.Printf("\nss%v\n", tokenDesc)
	println(token)
	//g.String(200, token1.UserName)
	g.String(200, "json", token)
	//g.ReturnJson(200, context.Context{ReturnStruct: map[string]interface{}{
	//	"code": 1,
	//	"msg":  "okk",
	//}})
}

func GetToken(g *gin.Context) {
	token := g.PostForm("token")
	println(token)
	//tokenDesc, _ := pkg.DecryptionToken(token)
	//fmt.Printf("\nss", tokenDesc.ExpiresAt)
	//println(token1.UserName)
	//g.String(200, token1.UserName)
	//ctx.ReturnJson(200, context.Context{ReturnStruct: map[string]interface{}{
	//	"code": 1,
	//	"msg":  "okk",
	//}})
}
