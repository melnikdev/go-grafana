package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/request"
)

type IMovieService interface {
	FindById(id string) (repository.Movie, error)
	Create(movie request.CreateMovieRequest) (string, error)
	Update(id string, movie request.UpdateMovieRequest) error
	Delete(id string) error
	GetTop5Movie() ([]repository.Movie, error)
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

func (s MovieService) FindById(id string) (repository.Movie, error) {
	return s.MovieRepository.FindById(id)
}

func (s MovieService) Create(movie request.CreateMovieRequest) (string, error) {
	err := s.Validate.Struct(movie)

	if err != nil {
		return "", err
	}

	m := repository.Movie{
		Title: movie.Title,
	}

	return s.MovieRepository.Create(m)
}

func (s MovieService) Update(id string, movie request.UpdateMovieRequest) error {
	err := s.Validate.Struct(movie)

	if err != nil {
		return err
	}

	m := repository.Movie{
		Title: movie.Title,
	}

	return s.MovieRepository.Update(id, m)
}

func (s MovieService) Delete(id string) error {
	return s.MovieRepository.Delete(id)
}

func (s MovieService) GetTop5Movie() ([]repository.Movie, error) {
	movies, err := s.MovieRepository.GetTop5Movie()

	if err != nil {
		return nil, err
	}

	return movies, nil
}
