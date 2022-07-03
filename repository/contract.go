package repository

import "movies-api/entity"

type MovieRepository interface {
	Create(entity.Movie) error
	Get() ([]entity.Movie, error)
	GetById(id int) (entity.Movie, error)
	Update(entity.Movie) error
	Delete(id int) error
}
