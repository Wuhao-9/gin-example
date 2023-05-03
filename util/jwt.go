package util

import (
	"github/Wuhao-9/go-gin-example/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

// 声明是一个结构体，它包含授权、用户和过期时间等信息。
// 用户自定义的声明
type Claims struct {
	jwt.StandardClaims        // 继承自jwt的标准声明
	Account            string `json:"account"`
	Passwd             string `json:"passwd"`
}

func GenerateToken(account, pwd string) (token string, err error) {
	now := time.Now()
	expireTime := now.Add(time.Hour * 3)
	// 创建一个声明结构体
	claim_obj := Claims{jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "gin-example-proj"},
		account, pwd,
	}
	// 创建一个Token
	// 使用此方法创建Token时，声明将自动添加到令牌中，无需手动添加。
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim_obj)
	// 使用服务器本地的密匙对Token进行签名，并返回签名后的Token字符串
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims方法用于解析 JWT 令牌
	// - tokenString：要解析的 JWT 令牌字符串。
	// - claims：一个实现了 jwt.Claims 接口的结构体，用于存储解析后的声明。
	// - keyFunc：一个函数，用于提供用于验证签名的密钥。
	// 返回值是一个 *jwt.Token 和一个error。如果解析成功，则返回一个包含令牌信息的 Token 对象；否则返回一个空指针和一个错误对象。
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(*jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
