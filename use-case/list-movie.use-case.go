package use_case

import (
	"movies-api/entity"
	"movies-api/repository"
)

type listMovie struct {
	MovieRepository repository.MovieRepository
}

type ListMovieUseCase interface {
	Execute() ([]entity.Movie, error)
}

func NewListMovieUseCase(r repository.MovieRepository) ListMovieUseCase {
	return &listMovie{r}
}

func (cm *listMovie) Execute() ([]entity.Movie, error) {
	data, err := cm.MovieRepository.Get()
	if err != nil {
		return nil, err
	}
	return data, nil
}
