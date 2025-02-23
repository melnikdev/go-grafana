package service

import "github.com/melnikdev/go-grafana/internal/repository"

type IMovieService interface {
	FindById(id string) (repository.Movie, error)
	Create(movie repository.Movie) (string, error)
}

type MovieService struct {
	MovieRepository repository.IMovieRepository
}

func NewMovieService(repo repository.IMovieRepository) *MovieService {
	return &MovieService{
		MovieRepository: repo,
	}
}

func (s MovieService) FindById(id string) (repository.Movie, error) {
	return s.MovieRepository.FindById(id)
}

func (s MovieService) Create(movie repository.Movie) (string, error) {
	return s.MovieRepository.Create(movie)
}
