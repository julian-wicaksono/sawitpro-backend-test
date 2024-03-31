package util

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userId int64, expiration time.Time, jwtKey []byte) (token string, err error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(int(userId)),
			ExpiresAt: expiration.Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = newToken.SignedString(jwtKey)
	return
}

func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
