package models

import (
	"errors"
	"time"

	db "github.com/Tavasiev/cws-backend/dbconn"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

const tokenExpiredTime = 1 // 1 - примерно 10-20 секунд, после токен просрочен
//const tokenExpiredTime = 1440
const refreshTokenExpiredMinutes = 201600

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

func (logResp *LoginResponse) newRefreshToken(userID int) error {

	newToken, err := uuid.NewV4()
	if err != nil {
		return err
	}

	logResp.UserID = userID
	logResp.RefreshToken = newToken.String()

	dur := time.Minute * time.Duration(refreshTokenExpiredMinutes)
	logResp.RefreshTokenExpiration = time.Now().Add(dur)

	err = logResp.saveTokenData(newToken.String())
	if err != nil {
		return err
	}
	return nil
}

// saveTokenData expired existing and create new token for user
func (logResp *LoginResponse) saveTokenData(uuid string) error {

	var sessNew Sessions

	sessNew.UserID = logResp.UserID
	sessNew.RefreshToken = uuid
	sessNew.RefreshTokenExpired = logResp.RefreshTokenExpiration

	_, err := db.Conn.Model(&sessNew).Returning("*").Insert()
	if err != nil {
		return errors.New("Ошибка сохранения новой сессии")
	}

	return nil
}
