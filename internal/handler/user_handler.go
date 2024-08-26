package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8000/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

type UserHandler struct {
	serv service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (*UserHandler) LoginWithGoogle(c *gin.Context) {
	googleOauthConfig.ClientID = os.Getenv("CLIENT_ID")
	googleOauthConfig.ClientSecret = os.Getenv("CLIENT_SECRET")

	oauthState := generateStateOauthCookie()
	authURL := googleOauthConfig.AuthCodeURL(oauthState)

	fmt.Println("url = " + authURL)

	// temporary with string
	c.String(http.StatusOK, authURL)
}

func generateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func (h *UserHandler) GetGoogleDetails(c *gin.Context) {
	token, err := googleOauthConfig.Exchange(context.Background(), c.Request.FormValue("code"))
	if err != nil {
		response.FailOrError(c, 500, "error occured when transfer authorization code into token", err)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		response.FailOrError(c, 500, "error occured when trying get access token", err)
		return
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		response.FailOrError(c, 500, "io error", err)
		return
	}

	var container dto.AuthenticatedUser
	err = json.Unmarshal(content, &container)
	if err != nil {
		response.FailOrError(c, 500, err.Error(), err)
		return
	}

	isExist, err := h.serv.EmailExists(container.Email)
	if err != nil {
		response.FailOrError(c, 500, "failed checking email", err)
	}
	if !isExist {
		h.serv.RegisterUserFromGoogle(container.Email)
	}

	data, err := h.serv.GenerateToken(container.Email)
	if err != nil {
		response.FailOrError(c, 500, "Failed generating token", err)
		return
	}
	response.Success(c, 200, "success", gin.H{"token": data})
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		response.ErrorEmptyField(c)
		return
	}
	isExist, err := h.serv.EmailExists(email)
	if err != nil {
		response.FailOrError(c, 500, "failed checking email", err)
		return
	}
	if isExist {
		response.FailOrError(c, 400, "email has been used", err)
		return
	}

	password := c.PostForm("password")
	if password == "" {
		response.ErrorEmptyField(c)
		return
	}
	password_konfirm := c.PostForm("password_konfirm")
	if password_konfirm != password {
		response.FailOrError(c, 400, "failed creating user", errors.New("konfirmasi password gagal"))
		return
	}

	phone := c.PostForm("phone")

	request := dto.UserCreateReq{
		Email:    email,
		Password: password,
		Phone:    phone,
	}

	user, err := h.serv.RegisterUser(request)
	if err != nil {
		response.FailOrError(c, 500, "Failed hasshing password", err)
		return
	}

	resp, err := h.serv.SendConfirmationEmail(user)
	if err != nil {
		response.FailOrError(c, 500, "Failed sending confirmation email", err)
		return
	}
	response.Success(c, 200, "Success sending confirmation email", resp)
}

func (h *UserHandler) CreateConfirmedUser(c *gin.Context) {
	tokenTemp := c.Query("token")
	claims := &dto.RegisterClaims{}
	token, err := jwt.ParseWithClaims(tokenTemp, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	input := dto.UserCreateReq{
		Email:    claims.Email,
		Password: claims.HPassword,
		Phone:    claims.Phone,
	}

	resp, err := h.serv.CreateUser(input)
	if err != nil {
		response.FailOrError(c, 500, "Failed creating user", err)
		return
	}
	response.Success(c, 201, "Success creating user", resp)
}

func (h *UserHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		response.ErrorEmptyField(c)
		return
	}
	password := c.PostForm("password")

	req := dto.UserLoginReq{
		Email:    email,
		Password: password,
	}

	data, err := h.serv.Login(req)
	if err != nil {
		response.FailOrError(c, 500, "Failed login", err)
		return
	}
	response.Success(c, 200, "Login succeed", gin.H{"token": data})
}
