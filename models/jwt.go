package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const tokenExpiredTime = 1 // 1 - примерно 10-20 секунд, после токен просрочен
//const tokenExpiredTime = 1440

// Sessions godoc
type Sessions struct {
	tableName           struct{}  `sql:"cws_sessions"`
	ID                  int       `json:"id" sql:",pk"`
	UserID              int       `json:"user_id"`
	RefreshToken        string    `json:"refresh_token" description:"Токен для обновления"`
	SessionEnd          time.Time `json:"session_end" description:"Дата когда токен отозван"`
	RefreshTokenUsed    time.Time `json:"refresh_token_used" description:"Дата использования токена"`
	RefreshTokenExpired time.Time `json:"refrash_expired" description:"Дата протухания токена"`
	CreatedAt           time.Time `sql:"default:now()" json:"created_at" description:"Дата создания"`
}

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
