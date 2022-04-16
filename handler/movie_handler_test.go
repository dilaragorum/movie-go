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

	t.Run("get movie error - ErrIDIsNotValid", func(t *testing.T) {
		req, _ := http.NewRequest("GET", reqURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovie(1).
			Return(model.Movie{}, service.ErrIDIsNotValid).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.GetMovie(rec, req, ps)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("get movie error - ErrMovieNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", reqURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovie(1).
			Return(model.Movie{}, service.ErrMovieNotFound).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.GetMovie(rec, req, ps)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("get movie error - InternalServerError", func(t *testing.T) {
		req, _ := http.NewRequest("GET", reqURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetMovie(1).
			Return(model.Movie{}, errors.New("")).
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
	t.Run("Error create movie - ErrTitleIsNotEmpty - Bad Request ", func(t *testing.T) {
		createdMovieReq := model.Movie{Title: ""}
		jsonStr, _ := json.Marshal(createdMovieReq)
		req, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			CreateMovie(model.Movie{Title: ""}).
			Return(service.ErrTitleIsNotEmpty).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.CreateMovie(rec, req, nil)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Error create movie - InternalServerError", func(t *testing.T) {
		createdMovieReq := model.Movie{Title: "Test Movie"}
		jsonStr, _ := json.Marshal(createdMovieReq)
		req, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			CreateMovie(model.Movie{Title: "Test Movie"}).
			Return(errors.New("")).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.CreateMovie(rec, req, nil)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("Success", func(t *testing.T) {
		createdMovieReq := model.Movie{Title: "Test Movie"}
		jsonStr, _ := json.Marshal(createdMovieReq)
		req, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			CreateMovie(model.Movie{Title: "Test Movie"}).
			Return(nil).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.CreateMovie(rec, req, nil)

		responseBodyStr := rec.Body.String()
		assert.Equal(t, "Movie is successfully created", responseBodyStr)
	})
}

func TestMovieHandler_DeleteMovie(t *testing.T) {
	movieID := "1"
	requestURL := fmt.Sprintf("/movies/%s", movieID)
	ps := httprouter.Params{
		{Key: "id", Value: movieID},
	}

	t.Run("delete movie error - Bad Request", func(t *testing.T) {
		type testCase struct {
			serviceErr error
			httpErr    int
		}

		testErrors := []testCase{
			{serviceErr: service.ErrIDIsNotValid, httpErr: http.StatusBadRequest},
			{serviceErr: service.ErrMovieNotFound, httpErr: http.StatusBadRequest},
		}

		for _, testError := range testErrors {
			req, _ := http.NewRequest(http.MethodDelete, requestURL, http.NoBody)
			rec := httptest.NewRecorder()

			mockService := service.NewMockIMovieService(gomock.NewController(t))
			mockService.
				EXPECT().
				DeleteMovie(1).
				Return(testError.serviceErr).
				Times(1)

			mh := NewMovieHandler(mockService)

			mh.DeleteMovie(rec, req, ps)

			assert.Equal(t, testError.httpErr, rec.Code)
		}

	})
	t.Run("delete movie error - Internal Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, requestURL, http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteMovie(1).
			Return(errors.New("")).
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
			Return(nil).
			Times(1)

		mh := NewMovieHandler(mockService)

		mh.DeleteMovie(rec, req, ps)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "Movie has been deleted succesfully", rec.Body.String())
	})
}

func TestMovieHandler_DeleteAllMovies(t *testing.T) {
	t.Run("delete all movie - Internal Server Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			DeleteAllMovie().
			Return(errors.New("")).
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
			Return(nil)

		mh := NewMovieHandler(mockService)
		mh.DeleteAllMovies(rec, req, nil)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "Movies successfully deleted", rec.Body.String())
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

	t.Run("update movie error - Bad Request", func(t *testing.T) {
		type testCase struct {
			returnedServiceErr     error
			expectedHTTPStatusCode int
		}

		testErrors := []testCase{
			{returnedServiceErr: service.ErrIDIsNotValid, expectedHTTPStatusCode: http.StatusBadRequest},
			{returnedServiceErr: service.ErrTitleIsNotEmpty, expectedHTTPStatusCode: http.StatusBadRequest},
		}

		for _, testError := range testErrors {
			req, _ := http.NewRequest(http.MethodPatch, requestURL, bytes.NewBuffer(jsonStr))
			rec := httptest.NewRecorder()

			mockService := service.NewMockIMovieService(gomock.NewController(t))
			mockService.
				EXPECT().
				UpdateMovie(1, updatedMovie).
				Return(testError.returnedServiceErr).
				Times(1)

			mh := NewMovieHandler(mockService)
			mh.UpdateMovie(rec, req, ps)

			assert.Equal(t, testError.expectedHTTPStatusCode, rec.Code)
		}
	})

	t.Run("update movie error - Status Not Found Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, requestURL, bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			UpdateMovie(1, updatedMovie).
			Return(service.ErrMovieNotFound).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.UpdateMovie(rec, req, ps)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("update movie error - Internal Server Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, requestURL, bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockService := service.NewMockIMovieService(gomock.NewController(t))
		mockService.
			EXPECT().
			UpdateMovie(1, updatedMovie).
			Return(errors.New("")).
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
			Return(nil).
			Times(1)

		mh := NewMovieHandler(mockService)
		mh.UpdateMovie(rec, req, ps)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "Successfully Updated", rec.Body.String())
	})
}
