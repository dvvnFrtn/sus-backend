package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticatedUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

type UserCreateReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	OauthID  string `json:"oauth_id"`
}

type UserCreateResp struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	ID   string `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
	jwt.RegisteredClaims
}

func NewUserClaims(id string, role string, exp time.Duration) UserClaims {
	return UserClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}

type RegisterClaims struct {
	Email     string `json:"email" binding:"required"`
	HPassword string `json:"h_password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	jwt.RegisteredClaims
}

func NewRegistrationClaims(email string, pass string, phone string, exp time.Duration) RegisterClaims {
	return RegisterClaims{
		Email:     email,
		HPassword: pass,
		Phone:     phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
