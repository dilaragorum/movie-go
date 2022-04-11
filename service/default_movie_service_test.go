package service

import (
	"github.com/dilaragorum/movie-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultMovieService_GetMovie(t *testing.T) {
	type testCase struct {
		id int
	}

	testCases := []testCase{
		{id: -2},
		{id: 0},
	}

	for _, test := range testCases {
		dms := NewDefaultMovieService(nil)
		_, err := dms.GetMovie(test.id)
		assert.ErrorIs(t, err, ErrIDIsNotValid)
	}
}

func TestDefaultMovieService_CreateMovie(t *testing.T) {
	dms := NewDefaultMovieService(nil)
	_, err := dms.CreateMovie(model.Movie{Title: ""})
	assert.ErrorIs(t, err, ErrTitleIsNotEmpty)
}

func TestDefaultMovieService_DeleteMovie(t *testing.T) {
	dms := NewDefaultMovieService(nil)
	_, err := dms.DeleteMovie(0)
	assert.ErrorIs(t, err, ErrIDIsNotValid)
}
