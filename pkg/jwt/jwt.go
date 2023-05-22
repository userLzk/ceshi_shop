package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var SigningKey = []byte("zidingyi")

//设置claims
type ClaimsCustom struct {
	CarryData //携带信息
	//Msg string `json:"msg"`
	jwt.StandardClaims
}

//定义jwt携带信息
type CarryData struct {
	UserName string `json:"UserName"`
	UserAge  int    `json:"UserAge"`
}

type JwtController struct {
}

//自定义结构存储
func NewCarryData(userName string, age int) CarryData {
	return CarryData{
		UserName: userName,
		UserAge:  age,
	}
}

//生成验证token
func EncryptionToken(Car CarryData) (token string) {

	//签证生产时间创建
	Iat := time.Now().Add(time.Second * 3600)

	newClaims := ClaimsCustom{
		CarryData: Car,
		StandardClaims: jwt.StandardClaims{
			Issuer: "test",
			//IssuedAt:  Iat, //签证生成时间
			ExpiresAt: Iat.Unix(), //签证有效期
		},
	}
	fmt.Sprintf("cl::", newClaims)
	//初始化token结构
	tokenStructure := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	//生成token
	token, err := tokenStructure.SignedString(SigningKey)
	if err != nil {
		panic(err)
	}
	return token
}

//验证token有效性
func DecryptionToken(signing string) (*ClaimsCustom, error) {

	token, err := jwt.ParseWithClaims(signing, &ClaimsCustom{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		fmt.Printf("err::%v", err)
	}
	if claims, ok := token.Claims.(*ClaimsCustom); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

//更新token信息
func RenewToken(claims ClaimsCustom) string {

	if withinLimit(claims.ExpiresAt, 6) {

		return EncryptionToken(claims.CarryData)
	}
	return ""
}

//检测是否满足续签条件
func withinLimit(s int64, l int64) bool {
	now := time.Now().Unix()
	return now-s < l
}
