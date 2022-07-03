package use_case

import (
	"movies-api/entity"
	"movies-api/repository"
)

type createMovie struct {
	MovieRepository repository.MovieRepository
}

type CreateMovieUseCase interface {
	Execute(m *entity.Movie) error
}

func NewCreateMovieUseCase(r repository.MovieRepository) CreateMovieUseCase {
	return &createMovie{r}
}

func (cm *createMovie) Execute(m *entity.Movie) error {
	err := cm.MovieRepository.Create(*m)
	if err != nil {
		return err
	}
	return nil
}
