package verfiy

import (
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
		//解析token信息
		claims, err := jwt.DecryptionToken(auth)

		if err != nil {
			//检测是否因为时间过期导致的 失效
			if strings.Contains(err.Error(), "expired") {
				//检测是否满足续签条件
				NewToken := jwt.RenewToken(*claims)
				//检测是否续签成功
				if NewToken != "" {
					//设置头信息
					c.Header("newToken", NewToken)
					c.Request.Header.Set("Authorization", NewToken)
					c.Next()
					return
				}
				//续签失败
				c.Abort()
				c.String(http.StatusOK, err.Error())
				return
			}
		}
		c.Next()
	}
}
