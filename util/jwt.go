package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shyptr/jianshu/setting"
	"strconv"
	"time"
)

func GeneraToken(id uint64, age time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: int64(age),
		Id:        fmt.Sprintf("%d", id),
	})
	return claims.SignedString([]byte(setting.GetJwtSecret()))
}

func ParseToken(token string) (uint64, error) {
	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(setting.GetJwtSecret()), nil
	})
	if err != nil {
		return 0, err
	}
	standardClaims := claims.Claims.(*jwt.StandardClaims)
	expiresAt := standardClaims.VerifyExpiresAt(time.Now().Unix(), true)
	if !expiresAt {
		return 0, errors.New("token失效")
	}
	return strconv.ParseUint(standardClaims.Id, 10, 0)
}
