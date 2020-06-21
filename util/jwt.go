package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shyptr/jianshu/setting"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	Id    int
	Root  bool
	State int
}

var ErrTokenExpire = errors.New("token过期")

func GeneraToken(id int, root bool, state int, age time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(age),
			Id:        fmt.Sprintf("%d", id),
		},
		Id:    id,
		Root:  root,
		State: state,
	})
	return claims.SignedString([]byte(setting.GetJwtSecret()))
}

func ParseToken(token string) (UserClaims, error) {
	claims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(setting.GetJwtSecret()), nil
	})
	if err != nil {
		return UserClaims{}, err
	}
	userClaims := claims.Claims.(*UserClaims)
	expiresAt := userClaims.VerifyExpiresAt(time.Now().Unix(), true)
	if !expiresAt {
		return UserClaims{}, ErrTokenExpire
	}
	return *userClaims, nil
}
