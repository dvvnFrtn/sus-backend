package jwt

import (
	"errors"
	"fmt"
	"os"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(payload sqlc.User) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 3
	}
	tokenJwtSementara := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.NewUserClaims(payload.ID, payload.Role, exp))
	tokenJwt, err := tokenJwtSementara.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}

func GenerateConfirmationToken(payload dto.UserCreateReq) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 3
	}
	tokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.NewRegistrationClaims(payload.Email, payload.Password, payload.Phone, exp))
	token, err := tokenTemp.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeToken(signedToken string) (*dto.UserClaims, error) {
	dcd, err := jwt.ParseWithClaims(signedToken, &dto.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("wrong signging method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token has been tampered with")
	}
	if !dcd.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims := dcd.Claims.(*dto.UserClaims)

	return claims, nil
}
