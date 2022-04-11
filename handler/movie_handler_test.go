package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/service"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieHandler_GetMovies(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovies().
			Return([]model.Movie{}, errors.New("oops!")).
			Times(1)

		mh := NewMovieHandler(mockService)

		mh.GetMovies(rec, req, nil)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovies().
			Return([]model.Movie{{ID: 1, Title: "Film"}}, nil).
			Times(1)

		mh := NewMovieHandler(mockService)

		mh.GetMovies(rec, req, nil)

		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedMovies []model.Movie
		json.NewDecoder(rec.Body).Decode(&returnedMovies)

		assert.Equal(t, 1, returnedMovies[0].ID)
		assert.Equal(t, "Film", returnedMovies[0].Title)
	})
}

func TestMovieHandler_GetMovie(t *testing.T) {
	movieID := "1"
	reqURL := fmt.Sprintf("/movies/%s", movieID)
	ps := httprouter.Params{
		{Key: "id", Value: movieID},
	}

	t.Run("get movie error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", reqURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovie(1).
			Return(model.Movie{}, errors.New("oops!")).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.GetMovie(rec, req, ps)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", reqURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovie(1).
			Return(model.Movie{ID: 1}, nil).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.GetMovie(rec, req, ps)

		assert.Equal(t, http.StatusOK, rec.Code)

		var ReturnMovie model.Movie
		json.NewDecoder(rec.Body).Decode(&ReturnMovie)

		assert.Equal(t, 1, ReturnMovie.ID)
	})
}

func TestMovieHandler_CreateMovie(t *testing.T) {
	createdMovieReq := model.Movie{Title: "Test Movie"}
	jsonStr, _ := json.Marshal(createdMovieReq)

	t.Run("create movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			CreateMovie(model.Movie{Title: "Test Movie"}).
			Return("", errors.New("Ups!")).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.CreateMovie(rec, req, nil)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			CreateMovie(model.Movie{Title: "Test Movie"}).
			Return("New movie is successfully added", nil).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.CreateMovie(rec, req, nil)

		responseBodyStr := rec.Body.String()
		assert.Equal(t, "New movie is successfully added", responseBodyStr)
	})
}

func TestMovieHandler_DeleteMovie(t *testing.T) {
	movieID := "1"
	requestURL := fmt.Sprintf("/movies/%s", movieID)
	ps := httprouter.Params{
		{Key: "id", Value: movieID},
	}

	t.Run("delete movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, requestURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteMovie(1).
			Return("", errors.New("Upps!")).
			Times(1)

		mh := NewMovieHandler(mockService)

		mh.DeleteMovie(rec, req, ps)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, requestURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteMovie(1).
			Return("Movie is successfully deleted", nil)

		mh := NewMovieHandler(mockService)

		mh.DeleteMovie(rec, req, ps)

		assert.Equal(t, "Movie is successfully deleted", rec.Body.String())
	})
}

func TestMovieHandler_DeleteAllMovies(t *testing.T) {
	t.Run("delete all movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteAllMovie().
			Return("", errors.New("Ops!")).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.DeleteAllMovies(rec, req, nil)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("delete all movie successfully", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteAllMovie().
			Return("All movies are successfully deleted", nil)

		mh := NewMovieHandler(mockService)
		mh.DeleteAllMovies(rec, req, nil)

		assert.Equal(t, "All movies are successfully deleted", rec.Body.String())
	})
}

func TestMovieHandler_UpdateMovie(t *testing.T) {
	movieID := "1"
	requestURL := fmt.Sprintf("/movies/%s", movieID)
	updatedMovie := model.Movie{Title: "Test Movie"}
	jsonStr, _ := json.Marshal(updatedMovie)
	ps := httprouter.Params{
		{Key: "id", Value: movieID},
	}

	t.Run("update movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, requestURL, bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			UpdateMovie(1, updatedMovie).
			Return("", errors.New("Ups!")).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.UpdateMovie(rec, req, ps)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("update movie successfully", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, requestURL, bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			UpdateMovie(1, updatedMovie).
			Return("Movie is successfully updated", nil).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.UpdateMovie(rec, req, ps)

		responseBodyStr := rec.Body.String()
		assert.Equal(t, "Movie is successfully updated", responseBodyStr)
	})
}
