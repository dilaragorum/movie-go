package service

import (
	"errors"
	"github.com/dilaragorum/movie-go/model"
)

type MockMovieService struct {
	ReturnErr bool
}

func (m MockMovieService) GetMovies() ([]model.Movie, error) {
	if m.ReturnErr {
		return []model.Movie{}, errors.New("problem")
	}

	movies := []model.Movie{
		{ID: 1},
		{ID: 2},
	}
	return movies, nil
}

func (m MockMovieService) GetMovie(id int) (model.Movie, error) {
	if m.ReturnErr {
		return model.Movie{}, errors.New("problem")
	}
	if id <= 0 {
		return model.Movie{}, ErrIDIsNotValid
	}
	return model.Movie{ID: 1}, nil
}

func (m MockMovieService) CreateMovie(title model.Movie) (message string, err error) {
	if m.ReturnErr {
		return "", errors.New("problem")
	}
	return "New movie is successfully added", nil
}

func (m MockMovieService) DeleteMovie(id int) (message string, err error) {

	if m.ReturnErr {
		return "", errors.New("problem")
	}
	return "Movie is successfully deleted", nil
}

func (m MockMovieService) DeleteAllMovie() (message string, err error) {
	if m.ReturnErr {
		return "", errors.New("problem")
	}

	return "All movies are successfully deleted", nil
}

func (m MockMovieService) UpdateMovie(id int, movie model.Movie) (message string, err error) {
	if m.ReturnErr {
		return "", errors.New("problem")
	}

	return "Movie is successfully updated", nil
}
