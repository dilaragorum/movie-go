package repository

import (
	"github.com/dilaragorum/movie-go/model"
)

type IMovieRepository interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) (string, error)
	DeleteMovie(id int) (string, error)
	DeleteAllMovies() (string, error)
	UpdateMovie(id int, movie model.Movie) (string, error)
}
