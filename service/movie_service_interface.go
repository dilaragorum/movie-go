package service

import "github.com/dilaragorum/movie-go/model"

// mockgen -source service/movie_service_interface.go -destination service/mock_movie_service.go -package service
type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) (string, error)
	DeleteMovie(id int) (string, error)
	DeleteAllMovie() (string, error)
	UpdateMovie(id int, movie model.Movie) (string, error)
}
