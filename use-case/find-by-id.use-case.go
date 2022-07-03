package use_case

import (
	"movies-api/entity"
	"movies-api/repository"
)

type findByIdMovie struct {
	MovieRepository repository.MovieRepository
}

type FindByIdMovieUseCase interface {
	Execute(id int) (entity.Movie, error)
}

func NewFindByIdMovieUseCase(r repository.MovieRepository) FindByIdMovieUseCase {
	return &findByIdMovie{r}
}

func (cm *findByIdMovie) Execute(id int) (entity.Movie, error) {
	data, err := cm.MovieRepository.GetById(id)
	if err != nil {
		return entity.Movie{}, err
	}
	return data, nil
}
