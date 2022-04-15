package repository

import (
	"errors"
	"github.com/dilaragorum/movie-go/model"
)

type inmemoryMovieRepository struct {
	Movies []model.Movie
}

func NewInMemoryMovieRepository() *inmemoryMovieRepository {
	var movies = []model.Movie{
		{ID: 1, Title: "Naruto"},
		{ID: 2, Title: "Closer"},
		{ID: 3, Title: "Mr Nobody"},
	}

	return &inmemoryMovieRepository{
		Movies: movies,
	}
}

func (i *inmemoryMovieRepository) GetMovies() ([]model.Movie, error) {
	return i.Movies, nil
}

func (i *inmemoryMovieRepository) GetMovie(id int) (model.Movie, error) {
	for _, movie := range i.Movies {
		if movie.ID == id {
			return movie, nil
		}
	}
	return model.Movie{}, errors.New("Movie cannot be found")
}

func (i *inmemoryMovieRepository) CreateMovie(movie model.Movie) (string, error) {
	movie.ID = len(i.Movies) + 1
	i.Movies = append(i.Movies, movie)

	return "New movie is successfully added", nil
}

func (i *inmemoryMovieRepository) DeleteMovie(id int) (string, error) {
	exist := false
	var newMovieList []model.Movie
	for _, movie := range i.Movies {
		if id == movie.ID {
			exist = true
		}
		if id != movie.ID {
			newMovieList = append(newMovieList, movie)
		}
	}

	if !exist {
		return "There is no movie with that id.", nil
	}

	i.Movies = newMovieList
	return "Movie is successfully deleted", nil
}

func (i *inmemoryMovieRepository) DeleteAllMovies() (string, error) {
	if i.Movies == nil {
		return "", errors.New("Movies have been already deleted")
	}
	i.Movies = nil
	return "All movies are successfully deleted", nil
}

func (i *inmemoryMovieRepository) UpdateMovie(id int, movie model.Movie) (string, error) {
	for k := 0; k < len(i.Movies); k++ {
		if i.Movies[k].ID == id {
			i.Movies[k].Title = movie.Title
		}
	}
	return "Movie is successfully updated", nil
}
