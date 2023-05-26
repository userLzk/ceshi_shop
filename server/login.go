package server

import (
	"net/http"

	"ceshi_shop/pkg/jwt"
	"github.com/gin-gonic/gin"
)

//创建对象

var LoginController loginController

type loginController struct {
}

func Login(g *gin.Context) {
	car := jwt.NewCarryData("zhagnsan", 17)
	accessToken, refreshToken := jwt.EncryptionToken(car)
	jwt.DecryptionToken(accessToken, refreshToken)
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "鉴权成功",
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func GetToken(g *gin.Context) {
	//logrus.WithFields()
	g.JSON(http.StatusOK, gin.H{"code": 200,
		"msg": "鉴权成功"})

}
