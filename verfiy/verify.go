package verfiy

import (
	"ceshi_shop/Connection"
)

//设置私钥
var jwtKey = Connection.ConfigData.JwtConfig.Sign

//生成jwt——token
func CreateJwtToken() {

}
