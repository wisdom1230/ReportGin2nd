package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"user" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

type MyClaims struct {
	Login
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 8 //设置过期时间

var privateKey = []byte("genome2022")

func ValidateUser(claims MyClaims) error {
	user := claims.Login.User
	password := claims.Login.Password
	if user == "genome" && password == "password" {
		return nil
	}
	return errors.New("user or password error")
}

func GenerateToken(login Login) (string, error) {
	claims := MyClaims{
		login,
		jwt.StandardClaims{
			ExpiresAt: int64(TokenExpireDuration),
			Issuer:    "genome2022",
			IssuedAt:  time.Now().Unix(),
			Subject:   "token",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(privateKey)
	return token, err
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { return privateKey, nil })
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

