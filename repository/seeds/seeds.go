package seeds

import "movies-api/entity"

func PopulateMovies(movies *[]entity.Movie) {
	*movies = append(*movies, entity.Movie{
		ID:          1,
		Title:       "Inception",
		Description: "Very nice",
		Director: &entity.Director{
			Name: "Jhon",
			Age:  54,
		},
	})

	*movies = append(*movies, entity.Movie{
		ID:          2,
		Title:       "Inception2",
		Description: "Very nice",
		Director: &entity.Director{
			Name: "Jhon",
			Age:  57,
		},
	})

	*movies = append(*movies, entity.Movie{
		ID:          3,
		Title:       "Inception3",
		Description: "Very nice",
		Director: &entity.Director{
			Name: "Jhon",
			Age:  59,
		},
	})
}
