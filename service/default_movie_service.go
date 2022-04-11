package service

import (
	"errors"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/repository"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrTitleIsNotEmpty = errors.New("Movie title cannot be empty")
)

type DefaultMovieService struct {
	movieRepo repository.IMovieRepository
}

func NewDefaultMovieService(mRepo repository.IMovieRepository) *DefaultMovieService {
	return &DefaultMovieService{
		movieRepo: mRepo,
	}
}

func (d *DefaultMovieService) GetMovies() ([]model.Movie, error) {
	return d.movieRepo.GetMovies()
}

func (d *DefaultMovieService) GetMovie(id int) (model.Movie, error) {
	// id hiç bir zaman 0 veya 0 dan küçük olamaz. --> Business Logic
	if id <= 0 {
		return model.Movie{}, ErrIDIsNotValid
	}
	return d.movieRepo.GetMovie(id)
}

func (d *DefaultMovieService) CreateMovie(movie model.Movie) (string, error) {
	if movie.Title == "" {
		return "", ErrTitleIsNotEmpty
	}
	//İşi repoya delege ediyorum
	return d.movieRepo.CreateMovie(movie)
}

func (d *DefaultMovieService) DeleteMovie(id int) (string, error) {
	if id <= 0 {
		return "", ErrIDIsNotValid
	}
	//İşi repoya delege ediyorum
	return d.movieRepo.DeleteMovie(id)
}

func (d *DefaultMovieService) DeleteAllMovie() (string, error) {
	return d.movieRepo.DeleteAllMovies()
}

func (d *DefaultMovieService) UpdateMovie(id int, movie model.Movie) (string, error) {
	if id <= 0 {
		return "", ErrIDIsNotValid
	}

	if movie.Title == "" {
		return "", ErrTitleIsNotEmpty
	}

	return d.movieRepo.UpdateMovie(id, movie)
}
