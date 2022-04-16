package service

import (
	"errors"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/repository"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrTitleIsNotEmpty = errors.New("Movie title cannot be empty")
	ErrMovieNotFound   = errors.New("the movie cannot be found")
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
	if id <= 0 {
		return model.Movie{}, ErrIDIsNotValid
	}
	movie, err := d.movieRepo.GetMovie(id)

	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return model.Movie{}, ErrMovieNotFound
		}
	}
	return movie, nil
}

func (d *DefaultMovieService) CreateMovie(movie model.Movie) error {
	if movie.Title == "" {
		return ErrTitleIsNotEmpty
	}
	return d.movieRepo.CreateMovie(movie)
}

func (d *DefaultMovieService) DeleteMovie(id int) error {
	if id <= 0 {
		return ErrIDIsNotValid
	}

	err := d.movieRepo.DeleteMovie(id)
	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return ErrMovieNotFound
		}
		return err
	}

	return nil
}

func (d *DefaultMovieService) DeleteAllMovie() error {
	return d.movieRepo.DeleteAllMovies()
}

func (d *DefaultMovieService) UpdateMovie(id int, movie model.Movie) error {
	if id <= 0 {
		return ErrIDIsNotValid
	}

	if movie.Title == "" {
		return ErrTitleIsNotEmpty
	}

	err := d.movieRepo.UpdateMovie(id, movie)
	if errors.Is(err, repository.ErrMovieNotFound) {
		return ErrMovieNotFound
	}

	return nil
}
