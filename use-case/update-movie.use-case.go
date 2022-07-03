package use_case

import (
	"movies-api/entity"
	"movies-api/repository"
)

type updateMovie struct {
	MovieRepository repository.MovieRepository
}

type UpdateMovieUseCase interface {
	Execute(m *entity.Movie) error
}

func NewUpdateMovieUseCase(r repository.MovieRepository) UpdateMovieUseCase {
	return &updateMovie{r}
}

func (cm *updateMovie) Execute(m *entity.Movie) error {
	err := cm.MovieRepository.Update(*m)
	if err != nil {
		return err
	}
	return nil
}
