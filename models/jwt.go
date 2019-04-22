package models

import (
	"log"
	"time"

	"github.com/Tavasiev/cws-backend/dbconn"

	"github.com/dgrijalva/jwt-go"
)

var db = dbconn.GetConnect()

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

type (
	// LoginRequest requested data when logging in
	LoginRequest struct {
		Phone    int    `json:"phone"`
		Password string `json:"password"`
	}

	// LoginRefreshRequest godoc
	LoginRefreshRequest struct {
		RefreshToken string `json:"refresh"`
	}

	// TokenClaim JWT token structure
	TokenClaim struct {
		Role   string `json:"role"`
		UserID int    `json:"user_id"`
		Phone  int    `json:"login"`
		jwt.StandardClaims
	}

	// LoginResponse responsed when requesting token
	LoginResponse struct {
		UserID                 int       `json:"user_id"`
		Token                  string    `json:"token"`
		RefreshToken           string    `json:"refresh_token"`
		RefreshTokenExpiration time.Time `json:"refresh_expiration"`
	}
)

func AuthenticateUser(data LoginRequest) (LoginResponse, error) {

	var oper Clients

	var login LoginResponse
	err := db.Insert(&oper)

	if err != nil {
		//return login, err
		panic(err)
	}
	log.Println("---------------------------------------------")
	log.Panicln(oper.Password)
	log.Println("---------------------------------------------")

	// Comparing the password with the hash
	/*err = bcrypt.CompareHashAndPassword([]byte(oper.Password), []byte(data.Password))
	if err != nil {
		return login, err
	}*/

	return login, nil
}

/*
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
}*/
