package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const tokenExpiredTime = 1

//const tokenExpiredTime = 1440

type JwtClaims struct {
	UserID int    `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

func CreateJwtToken() (string, error) {
	claims := JwtClaims{
		2,
		"79888794725",
		jwt.StandardClaims{},
	}

	claims.IssuedAt = time.Now().Unix()
	dur := time.Minute * time.Duration(tokenExpiredTime)
	claims.ExpiresAt = time.Now().Add(dur).Unix()

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))

	if err != nil {
		return "", err
	}
	return token, nil
}
