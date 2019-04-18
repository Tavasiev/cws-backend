package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

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

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
