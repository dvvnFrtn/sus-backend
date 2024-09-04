package service

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
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
	SendConfirmationEmail(dto.UserCreateReq) (dto.UserCreateResp, error)
	RegisterUserFromGoogle(string) (*dto.ResponseID, error)
	CreateUser(dto.UserCreateReq) (*dto.ResponseID, error)
	Login(dto.UserLoginReq) (string, error)
	GenerateToken(string) (string, error)
	FindUserByID(string) (*dto.UserResponse, error)
	UpdateUser(string, dto.UserUpdateReq) (*dto.UserUpdateReq, error)
	AddUserCategory(string, []string) error
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

func (s *userService) SendConfirmationEmail(arg dto.UserCreateReq) (dto.UserCreateResp, error) {
	domain := strings.Split(arg.Email, "@")[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return dto.UserCreateResp{}, err
	}

	token, err := jwt.GenerateConfirmationToken(arg)
	if err != nil {
		return dto.UserCreateResp{}, err
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
		return dto.UserCreateResp{}, err
	}

	d := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	err = d.DialAndSend(m)
	if err != nil {
		return dto.UserCreateResp{}, err
	}

	resp := dto.UserCreateResp{
		Email: arg.Email,
		Phone: arg.Phone,
	}
	return resp, nil
}

func (s *userService) RegisterUserFromGoogle(email string) (*dto.ResponseID, error) {
	input := dto.UserCreateReq{
		Email:    email,
		Password: "",
		Phone:    "",
	}
	return s.CreateUser(input)
}

func (s *userService) CreateUser(arg dto.UserCreateReq) (*dto.ResponseID, error) {
	n, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}
	user := sqlc.AddUserParams{
		ID:          uuid.New().String(),
		Email:       arg.Email,
		Password:    sql.NullString{String: arg.Password, Valid: arg.Password != ""},
		OauthID:     sql.NullString{String: arg.OauthID, Valid: arg.OauthID != ""},
		Phone:       sql.NullString{String: arg.Phone, Valid: arg.Phone != ""},
		Name:        "User" + n.String(),
		Role:        "user",
		Address:     sql.NullString{},
		Img:         sql.NullString{},
		IsPremium:   sql.NullBool{Bool: false, Valid: true},
		Lvl:         sql.NullInt32{Int32: 1, Valid: true},
		Dob:         sql.NullTime{},
		Institution: sql.NullString{},
		CreatedAt:   sql.NullTime{Time: time.Now()},
		UpdatedAt:   sql.NullTime{Time: time.Now()},
	}

	_, err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseID{ID: user.ID}, nil
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

func (s *userService) FindUserByID(id string) (*dto.UserResponse, error) {
	user, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToUserResponse(&user), nil
}

func (s *userService) UpdateUser(id string, arg dto.UserUpdateReq) (*dto.UserUpdateReq, error) {
	dob, err := time.Parse("2006-01-02", arg.DOB)
	if err != nil {
		return nil, err
	}
	input := sqlc.UpdateUserByIDParams{
		Username:    sql.NullString{String: arg.Username, Valid: arg.Username != ""},
		Name:        arg.Name,
		Address:     sql.NullString{String: arg.Address, Valid: arg.Address != ""},
		Dob:         sql.NullTime{Time: dob, Valid: !dob.IsZero()},
		Institution: sql.NullString{String: arg.Institution, Valid: arg.Institution != ""},
		ID:          id,
	}
	_, err = s.repo.UpdateUser(input)
	if err != nil {
		return nil, err
	}
	return &arg, nil
}

func (s *userService) AddUserCategory(u_id string, cat_ids []string) error {
	for _, cat_id := range cat_ids {
		check := sqlc.UserCategoryExistsParams{
			CategoryID: cat_id,
			UserID:     u_id,
		}
		count, err := s.repo.UserCategoryExists(check)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		input := sqlc.CreateUserCategoryParams(check)

		_, err = s.repo.AddUserCategory(input)
		if err != nil {
			return err
		}
	}
	return nil
}
