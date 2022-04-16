package service

import (
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultMovieService_GetMovie(t *testing.T) {
	t.Run("Error getMovie - ErrIDIsNotValid", func(t *testing.T) {
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
	})
	t.Run("Error getMovie - ErrMovieNotFound", func(t *testing.T) {
		mockRepository := repository.NewMockIMovieRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			GetMovie(6).
			Return(model.Movie{}, repository.ErrMovieNotFound).
			Times(1)

		dms := NewDefaultMovieService(mockRepository)
		_, err := dms.GetMovie(6)

		assert.ErrorIs(t, err, ErrMovieNotFound)

	})

}

func TestDefaultMovieService_CreateMovie(t *testing.T) {
	t.Run("Error Create Movie - ErrTitleIsNotEmpty", func(t *testing.T) {
		dms := NewDefaultMovieService(nil)
		err := dms.CreateMovie(model.Movie{Title: ""})
		assert.ErrorIs(t, err, ErrTitleIsNotEmpty)
	})
	t.Run("Success Create Movie", func(t *testing.T) {
		movie := model.Movie{Title: "Test Movie"}
		mockRepository := repository.NewMockIMovieRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().CreateMovie(movie).
			Return(nil).
			Times(1)

		ms := NewDefaultMovieService(mockRepository)
		err := ms.CreateMovie(movie)

		assert.Nil(t, err)
	})

}

func TestDefaultMovieService_DeleteMovie(t *testing.T) {
	t.Run("Error Delete Movie - ErrIDIsNotValid", func(t *testing.T) {
		dms := NewDefaultMovieService(nil)
		err := dms.DeleteMovie(0)
		assert.ErrorIs(t, err, ErrIDIsNotValid)
	})
	t.Run("Error Delete Movie - ErrMovieNotFound", func(t *testing.T) {
		mockRepository := repository.NewMockIMovieRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			DeleteMovie(6).
			Return(repository.ErrMovieNotFound).
			Times(1)

		ms := NewDefaultMovieService(mockRepository)
		err := ms.DeleteMovie(6)
		assert.ErrorIs(t, err, ErrMovieNotFound)
	})

}

func TestDefaultMovieService_UpdateMovie(t *testing.T) {
	t.Run("Error Update Movie - IDIsNotValid", func(t *testing.T) {
		ms := NewDefaultMovieService(nil)
		err := ms.UpdateMovie(0, model.Movie{Title: ""})
		assert.ErrorIs(t, err, ErrIDIsNotValid)
	})
	t.Run("Error Update Movie - ErrTitleIsNotEmpty", func(t *testing.T) {
		ms := NewDefaultMovieService(nil)
		err := ms.UpdateMovie(3, model.Movie{Title: ""})
		assert.ErrorIs(t, err, ErrTitleIsNotEmpty)
	})
	t.Run("Error Update Movie - ErrMovieNotFound", func(t *testing.T) {
		movie := model.Movie{Title: "Test Movie"}
		mockRepository := repository.NewMockIMovieRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			UpdateMovie(6, movie).
			Return(repository.ErrMovieNotFound).
			Times(1)

		ms := NewDefaultMovieService(mockRepository)
		err := ms.UpdateMovie(6, movie)

		assert.ErrorIs(t, err, ErrMovieNotFound)
	})

	t.Run("Error Update Movie - ErrMovieNotFound", func(t *testing.T) {
		movie := model.Movie{Title: "Test Movie"}
		mockRepository := repository.NewMockIMovieRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			UpdateMovie(2, movie).
			Return(nil).
			Times(1)

		ms := NewDefaultMovieService(mockRepository)
		err := ms.UpdateMovie(2, movie)

		assert.Nil(t, err)
	})
}
