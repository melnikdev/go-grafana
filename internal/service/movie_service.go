package service

import "github.com/melnikdev/go-grafana/internal/repository"

type IMovieService interface {
	FindById() (repository.Restaurant, error)
}

type MovieService struct {
	MovieRepository repository.IMovieRepository
}

func NewMovieService(repo repository.IMovieRepository) *MovieService {
	return &MovieService{
		MovieRepository: repo,
	}
}

func (s MovieService) FindById() (repository.Restaurant, error) {
	return s.MovieRepository.FindById()
}
