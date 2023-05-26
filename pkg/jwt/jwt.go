package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

//刷新refresh token key
var refreshSecretKey = []byte("zidingyi")

// access token key
var accessSecretKey = []byte("zidingyi")

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
func EncryptionToken(Car CarryData) (string, string) {

	//创年token有效期
	Iat := time.Now().Add(time.Second * 30)
	//创建access有效期
	refreshTime := time.Now().Add(time.Second * 3600)
	accessClaims := ClaimsCustom{
		CarryData: Car,
		StandardClaims: jwt.StandardClaims{
			Issuer: "test",
			//IssuedAt:  Iat, //签证生成时间
			ExpiresAt: Iat.Unix(), //签证有效期
		},
	}
	//刷新时间token结构
	refreshClaims := ClaimsCustom{
		CarryData: Car,
		StandardClaims: jwt.StandardClaims{
			Issuer: "test",
			//IssuedAt:  Iat, //签证生成时间
			ExpiresAt: refreshTime.Unix(), //签证有效期
		},
	}
	//初始化token结构
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	//初始化token结构
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	//create refreshToken
	accessTokenSigned, err := accessToken.SignedString(accessSecretKey)
	if err != nil {
		panic(err)
	}
	//create refreshToken
	refreshTokenSigned, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		panic(err)
	}
	return accessTokenSigned, refreshTokenSigned
}

//验证token有效性
func DecryptionToken(accessTokenString, refreshTokenString string) (*ClaimsCustom, bool, error) {

	//解密accessToken
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &ClaimsCustom{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecretKey, nil
	})
	//检测是否需要刷新taoken@TODO 如果解密成功可以正确获取信息 无需花心token
	if claims, ok := accessToken.Claims.(*ClaimsCustom); ok && accessToken.Valid {
		return claims, false, nil
	}
	//解密refreshToken
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &ClaimsCustom{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})
	//解密失敗
	if err != nil {
		return nil, false, err
	}
	//检测是否需要刷新taoken
	if claims, ok := refreshToken.Claims.(*ClaimsCustom); ok && refreshToken.Valid {
		return claims, true, nil
	}
	return nil, false, errors.New("token error")
}

//更新token信息
//func RenewToken(claims ClaimsCustom) string {

//if withinLimit(claims.ExpiresAt, 10) {
//
//	return EncryptionToken(claims.CarryData)
//}
//return ""
//}

//检测是否满足续签条件
func withinLimit(s int64, l int64) bool {
	now := time.Now().Unix()
	return now-s < l
}
