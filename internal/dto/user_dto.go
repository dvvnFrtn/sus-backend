package dto

import (
	"sus-backend/internal/db/sqlc"
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

type UserUpdateReq struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	DOB         string `json:"dob"`
	Institution string `json:"institution"`
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

type UserResponse struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	OauthID     string    `json:"oauth_id"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Img         string    `json:"img"`
	IsPremium   bool      `json:"is_premium"`
	Lvl         int       `json:"lvl"`
	Dob         time.Time `json:"dob"`
	Institution string    `json:"institution"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToUserResponse(u *sqlc.User) *UserResponse {
	return &UserResponse{
		ID:          u.ID,
		Email:       u.Email,
		OauthID:     u.OauthID.String,
		Username:    u.Username.String,
		Name:        u.Name,
		Role:        u.Role,
		Phone:       u.Phone.String,
		Img:         u.Img.String,
		IsPremium:   u.IsPremium.Bool,
		Lvl:         int(u.Lvl.Int32),
		Dob:         u.Dob.Time,
		Institution: u.Institution.String,
		CreatedAt:   u.CreatedAt.Time,
		UpdatedAt:   u.UpdatedAt.Time,
	}
}
