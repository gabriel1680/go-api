package repository

import (
	"errors"
	"math/rand"
	"movies-api/entity"
	"movies-api/repository/seeds"
)

type InMemoryMovieRepository struct {
}

var Movies []entity.Movie

func init() {
	seeds.PopulateMovies(&Movies)
}

func NewMovieRepository() MovieRepository {
	return &InMemoryMovieRepository{}
}

func (mr *InMemoryMovieRepository) Create(movie entity.Movie) error {
	movie.ID = rand.Intn(100)
	Movies = append(Movies, movie)
	return nil
}

func (mr *InMemoryMovieRepository) Update(movie entity.Movie) error {
	for index, item := range Movies {
		if item.ID == movie.ID {
			Movies[index] = movie
			return nil
		}
	}
	return errors.New("movie with id not found")
}

func (mr *InMemoryMovieRepository) GetById(id int) (entity.Movie, error) {
	for _, item := range Movies {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Movie{}, errors.New("movie with id not found")
}

func (mr *InMemoryMovieRepository) Get() ([]entity.Movie, error) {
	return Movies, nil
}

func (mr *InMemoryMovieRepository) Delete(id int) error {
	for index, item := range Movies {
		if item.ID == id {
			Movies = append(Movies[:index], Movies[index+1:]...)
			return nil
		}
	}
	return errors.New("movie with id not found")
}
