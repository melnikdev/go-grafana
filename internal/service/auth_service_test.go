package service_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/melnikdev/go-grafana/internal/request"
	"github.com/melnikdev/go-grafana/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

// Имитация метода FindById
func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

// Имитация метода Create
func (m *MockUserRepository) Create(user model.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupAuthTest() (*service.AuthService, *MockUserRepository) {
	mockRepo := new(MockUserRepository)
	validate := validator.New()
	return service.NewAuthService(mockRepo, validate), mockRepo
}

func TestRegister(t *testing.T) {
	authService, mockRepo := setupAuthTest()

	validRequest := request.RegisterUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.On("Create", mock.Anything).Return("123", nil)

	userID, err := authService.Register(validRequest)

	assert.NoError(t, err)
	assert.Equal(t, "123", userID)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	authService, mockRepo := setupAuthTest()

	objectID, err := primitive.ObjectIDFromHex("123")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	storedUser := model.User{
		ID:       objectID,
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	validRequest := request.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.On("FindByEmail", validRequest.Email).Return(&storedUser, nil)

	token, err := authService.Login(validRequest)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLoginInvalidPassword(t *testing.T) {
	authService, mockRepo := setupAuthTest()

	objectID, err := primitive.ObjectIDFromHex("123")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)
	storedUser := model.User{
		ID:       objectID,
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	invalidRequest := request.LoginUserRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	mockRepo.On("FindByEmail", invalidRequest.Email).Return(&storedUser, nil)

	token, err := authService.Login(invalidRequest)

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
