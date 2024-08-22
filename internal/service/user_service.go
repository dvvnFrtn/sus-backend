package service

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	"sus-backend/pkg/bcrypt"
	"sus-backend/pkg/jwt"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/google/uuid"
)

type UserService interface {
	EmailExists(string) (bool, error)
	RegisterUser(dto.UserCreateReq) (dto.UserCreateReq, error)
	SendConfirmationEmail(dto.UserCreateReq) error
	RegisterUserFromGoogle(dto.UserCreateReq) (sql.Result, error)
	CreateUser(dto.UserCreateReq) (sql.Result, error)
	Login(dto.UserLoginReq) (string, error)
	GenerateToken(string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) EmailExists(email string) (bool, error) {
	count, err := s.repo.EmailExists(email)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *userService) RegisterUser(arg dto.UserCreateReq) (dto.UserCreateReq, error) {
	hashedPassword, err := bcrypt.HashValue(arg.Password)
	if err != nil {
		return dto.UserCreateReq{}, err
	}

	input := dto.UserCreateReq{
		Email:    arg.Email,
		Password: hashedPassword,
		Phone:    arg.Phone,
	}
	return input, err
}

func (s *userService) SendConfirmationEmail(arg dto.UserCreateReq) error {
	domain := strings.Split(arg.Email, "@")[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return err
	}

	token, err := jwt.GenerateConfirmationToken(arg)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("%s/account-confirm?token=%s", os.Getenv("BASE_URL"), token)

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromEmail := os.Getenv("FROM_EMAIL")
	fromName := os.Getenv("FROM_NAME")

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, fromEmail))
	m.SetHeader("To", arg.Email)
	m.SetHeader("Subject", "Email Confirmation")
	m.SetBody("text/plain", fmt.Sprintf("Please confirm your email by clicking on the following link: %s", link))

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) RegisterUserFromGoogle(arg dto.UserCreateReq) (sql.Result, error) {
	input := dto.UserCreateReq{
		Email:    arg.Email,
		Password: arg.Password,
		Phone:    arg.Phone,
	}
	return s.CreateUser(input)
}

func (s *userService) CreateUser(arg dto.UserCreateReq) (sql.Result, error) {
	user := sqlc.AddUserParams{
		ID:          uuid.New().String(),
		Email:       arg.Email,
		Password:    sql.NullString{String: arg.Password, Valid: arg.Password != ""},
		OauthID:     sql.NullString{String: arg.OauthID, Valid: arg.OauthID != ""},
		Phone:       sql.NullString{String: arg.Phone, Valid: arg.Phone != ""},
		Name:        "User",
		Role:        "user",
		Img:         sql.NullString{},
		IsPremium:   sql.NullBool{Bool: false, Valid: true},
		Lvl:         sql.NullInt32{Int32: 1, Valid: true},
		Dob:         sql.NullTime{},
		Institution: sql.NullString{},
		CreatedAt:   sql.NullTime{Time: time.Now()},
		UpdatedAt:   sql.NullTime{Time: time.Now()},
	}
	return s.repo.CreateUser(user)
}

func (s *userService) Login(arg dto.UserLoginReq) (string, error) {
	user, err := s.repo.FindByEmail(arg.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.ValidateHash(arg.Password, user.Password.String)
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) GenerateToken(email string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
