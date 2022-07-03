package use_case

import (
	"movies-api/repository"
)

type deleteMovie struct {
	MovieRepository repository.MovieRepository
}

type DeleteMovieUseCase interface {
	Execute(id int) error
}

func NewDeleteMovieUseCase(r repository.MovieRepository) DeleteMovieUseCase {
	return &deleteMovie{r}
}

func (cm *deleteMovie) Execute(id int) error {
	err := cm.MovieRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
