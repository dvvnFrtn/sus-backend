package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
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

	// isExist, err := h.service.EmailExists(container.Email)
	// if err != nil {
	// 	response.FailOrError(c, 500, "failed checking email", err)
	// }
	// if !isExist {
	// 	input := dto.UserCreateReq{
	// 		Email:    container.Email,
	// 		Password: "",
	// 		Name:     container.Name,
	// 	}
	// 	h.service.AddUser(input)
	// }

	response.Success(c, 200, "success", container)
}
