package util

import (
	"github.com/golang-jwt/jwt/v5"
)

var key = "miyuekey"

// jwt 携带的数据
type JwtPayload struct {
	Name   string `json:"name"`
	Age    int `json:"age"`
	Gender int `json:"gender"`
	// 內嵌 jwt 類型，用來验证 jwt 的有效性
	jwt.RegisteredClaims
}

// 签名
func Sign(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	sign, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return sign, err
}

// 验证签名
func Verify(sign string,payload *JwtPayload) error {
	_, err := jwt.ParseWithClaims(sign, payload, func(t *jwt.Token) (any, error) {
		return []byte(key), nil
	})
	return err
}
