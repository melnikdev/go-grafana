package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/request"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Create(r request.RegisterUserRequest) (string, error)
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

func (s AuthService) Create(r request.RegisterUserRequest) (string, error) {
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
