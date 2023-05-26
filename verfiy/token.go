package verfiy

import (
	"fmt"
	"net/http"
	"strings"

	"ceshi_shop/pkg/jwt"
	"github.com/gin-gonic/gin"
)

//检测token是否过期
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		auth := c.Request.Header.Get("Authorization")
		//检测是否有授权码
		if len(auth) == 0 {
			c.Abort()
			c.String(http.StatusOK, "未授权请先登录")
			return
		}
		//获取token信息
		tokenAuth := strings.Split(auth, " ")
		if tokenAuth[0] != "Bearer" {
			fmt.Printf("authorization string err!!")
		}
		//解析token信息
		claims, isUpdate, err := jwt.DecryptionToken(tokenAuth[1], tokenAuth[2])

		if err != nil {
			// 检测授权是否失效
			c.String(http.StatusOK, "授权失效请重新登录", err)
			c.Abort()
			return
		}
		//检测是否需要刷新 如果isUpdate为true的话整证明访问token( accessToken 已经失效了) refreshToken还未失效 这个时候需要刷新双token信息
		if isUpdate {
			tokenAuth[1], tokenAuth[2] = jwt.EncryptionToken(claims.CarryData)
			NewToken := strings.Join(tokenAuth, " ")
			c.Header("nowToken", NewToken)
			c.Request.Header.Set("Authorization", NewToken)
			//c.JSON(200, gin.H{
			//	"code": 200,
			//	"msg":  "鉴权成功",
			//	"data": gin.H{
			//		"accessToken":  tokenAuth[1],
			//		"refreshToken": tokenAuth[2],
			//	},
			//})
		}
		c.Next()
	}
}
