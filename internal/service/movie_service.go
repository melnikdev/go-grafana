package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-grafana/internal/model"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/request"
)

type IMovieService interface {
	FindById(id string) (*model.Movie, error)
	Create(r request.CreateMovieRequest) (string, error)
	Update(id string, r request.UpdateMovieRequest) error
	Delete(id string) error
	GetTopMovies(limit int64) ([]model.Movie, error)
}

type MovieService struct {
	MovieRepository repository.IMovieRepository
	Validate        *validator.Validate
}

func NewMovieService(repo repository.IMovieRepository, val *validator.Validate) *MovieService {
	return &MovieService{
		MovieRepository: repo,
		Validate:        val,
	}
}

func (s MovieService) FindById(id string) (*model.Movie, error) {
	return s.MovieRepository.FindById(id)
}

func (s MovieService) Create(r request.CreateMovieRequest) (string, error) {
	err := s.Validate.Struct(r)

	if err != nil {
		return "", err
	}

	m := model.Movie{
		Title: r.Title,
	}

	return s.MovieRepository.Create(m)
}

func (s MovieService) Update(id string, r request.UpdateMovieRequest) error {
	err := s.Validate.Struct(r)

	if err != nil {
		return err
	}

	m := model.Movie{
		Title: r.Title,
	}

	return s.MovieRepository.Update(id, &m)
}

func (s MovieService) Delete(id string) error {
	return s.MovieRepository.Delete(id)
}

func (s MovieService) GetTopMovies(limit int64) ([]model.Movie, error) {
	movies, err := s.MovieRepository.GetTopMovies(limit)

	if err != nil {
		return nil, err
	}

	return movies, nil
}
