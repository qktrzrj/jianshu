package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shyptr/hello-world-web/setting"
	"time"
)

var secret = []byte(setting.JWT_SECRET)

func GenToken(id int64, ages ...int) (string, error) {
	age := 1
	if len(ages) > 0 {
		age = ages[0]
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", id),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(24*age)).Unix(),
	})
	return token.SignedString(secret)
}

func ParseAnValidateToken(jwtToken string) (jwt.StandardClaims, error) {
	// 解析 token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 进行 alg 即签名算法校验
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return jwt.StandardClaims{}, err
	}

	// 校验有效性，并获取 Claims 中的值
	if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return jwt.StandardClaims{}, errors.New("非法token")
}
