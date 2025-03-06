package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/request"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(r request.RegisterUserRequest) (string, error)
	Login(r request.LoginUserRequest) (string, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
	Validate       *validator.Validate
}

func NewAuthService(repo repository.IUserRepository, val *validator.Validate) *AuthService {
	return &AuthService{
		UserRepository: repo,
		Validate:       val,
	}
}

func (s AuthService) Register(r request.RegisterUserRequest) (string, error) {
	err := s.Validate.Struct(r)

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user := model.User{
		Email:    r.Email,
		Password: string(hashedPassword),
	}

	return s.UserRepository.Create(user)
}

func (s AuthService) Login(r request.LoginUserRequest) (string, error) {
	var jwtKey = []byte("go_test_secret_key")

	err := s.Validate.Struct(r)

	if err != nil {
		return "", err
	}

	storedUser, err := s.UserRepository.FindByEmail(r.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(r.Password))

	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   storedUser.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
