package service

import "github.com/dilaragorum/movie-go/model"

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) (string, error)
	DeleteMovie(id int) (string, error)
	DeleteAllMovie() (string, error)
	UpdateMovie(id int, movie model.Movie) (string, error)
}
