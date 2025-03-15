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
)

type MockMovieRepository struct {
	mock.Mock
}

// Имитация метода FindById
func (m *MockMovieRepository) FindById(id string) (*model.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Movie), args.Error(1)
}

// Имитация метода Create
func (m *MockMovieRepository) Create(movie model.Movie) (string, error) {
	args := m.Called(movie)
	return args.String(0), args.Error(1)
}

// Имитация метода Update
func (m *MockMovieRepository) Update(id string, movie *model.Movie) error {
	args := m.Called(id, movie)
	return args.Error(0)
}

// Имитация метода Delete
func (m *MockMovieRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Имитация метода GetTopMovies
func (m *MockMovieRepository) GetTopMovies(limit int64) ([]model.Movie, error) {
	args := m.Called(limit)
	return args.Get(0).([]model.Movie), args.Error(1)
}

func setupTest() (*service.MovieService, *MockMovieRepository) {
	mockRepo := new(MockMovieRepository)
	validate := validator.New()
	return service.NewMovieService(mockRepo, validate), mockRepo
}

func TestCreateMovie(t *testing.T) {
	svc, mockRepo := setupTest()
	objectID, err := primitive.ObjectIDFromHex("123")
	movie := model.Movie{ID: objectID, Title: "Test Movie"}

	strId := movie.ID.Hex()
	mockRepo.On("Create", mock.Anything).Return(strId, nil)

	id, err := svc.Create(request.CreateMovieRequest{Title: "Test Movie"})
	assert.NoError(t, err)
	assert.Equal(t, strId, id)
}

func TestFindById(t *testing.T) {
	svc, mockRepo := setupTest()
	movie := model.Movie{Title: "Test Movie"}
	mockRepo.On("FindById", "123").Return(&movie, nil)

	fetchedMovie, err := svc.FindById("123")
	assert.NoError(t, err)
	assert.Equal(t, "Test Movie", fetchedMovie.Title)
}

func TestGetTopMovies(t *testing.T) {
	svc, mockRepo := setupTest()
	movie := model.Movie{Title: "Test Movie"}
	mockRepo.On("GetTopMovies", int64(1)).Return([]model.Movie{movie}, nil)

	topMovies, err := svc.GetTopMovies(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, topMovies)
}

func TestUpdateMovie(t *testing.T) {
	svc, mockRepo := setupTest()
	mockRepo.On("Update", "123", mock.Anything).Return(nil)

	err := svc.Update("123", request.UpdateMovieRequest{Id: 123, Title: "Updated Movie"})
	assert.NoError(t, err)
}

func TestDeleteMovie(t *testing.T) {
	svc, mockRepo := setupTest()
	mockRepo.On("Delete", "123").Return(nil)

	err := svc.Delete("123")
	assert.NoError(t, err)
}
