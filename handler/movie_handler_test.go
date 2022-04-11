package handler

import (
	"bytes"
	"encoding/json"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Get
func TestMovieHandler_GetMovies_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "POST not ALLOWED",
			reqMethod: http.MethodPost,
		},
		{
			name:      "PUT not ALLOWED",
			reqMethod: http.MethodPut,
		},
		{
			name:      "DELETE not ALLOWED",
			reqMethod: http.MethodDelete,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "/movies", http.NoBody)
			rec := httptest.NewRecorder()
			mh := NewMovieHandler(nil)
			mh.GetMovies(rec, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
		})
	}
}

func TestMovieHandler_GetMovies(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)

		mh.GetMovies(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/movies", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)

		mh.GetMovies(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedMovies []model.Movie
		json.NewDecoder(rec.Body).Decode(&returnedMovies)

		assert.Equal(t, 1, returnedMovies[0].ID)
		assert.Equal(t, 2, returnedMovies[1].ID)

	})
}

func TestMovieHandler_GetMovie_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "POST not allowed",
			reqMethod: http.MethodPost,
		},
		{
			name:      "PUT not allowed",
			reqMethod: http.MethodPut,
		},
		{
			name:      "DELETE not allowed",
			reqMethod: http.MethodDelete,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "/movie", http.NoBody)
			rec := httptest.NewRecorder()

			mh := &movieHandler{nil}
			mh.GetMovies(rec, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
		})
	}
}

func TestMovieHandler_GetMovie(t *testing.T) {
	t.Run("get movie error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/movie", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)
		mh.GetMovie(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/movie?id=1", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)

		mh.GetMovie(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var ReturnMovie model.Movie
		json.NewDecoder(rec.Body).Decode(&ReturnMovie)

		assert.Equal(t, 1, ReturnMovie.ID)

	})
}

func TestMovieHandler_CreateMovie_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "GET not allowed",
			reqMethod: http.MethodGet,
		},
		{
			name:      "PUT not allowed",
			reqMethod: http.MethodPut,
		},
		{
			name:      "DELETE not allowed",
			reqMethod: http.MethodDelete,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "/moviecreate", http.NoBody)
			rec := httptest.NewRecorder()

			mh := NewMovieHandler(nil)
			mh.CreateMovie(rec, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
		})
	}
}

func TestMovieHandler_CreateMovie(t *testing.T) {
	t.Run("create movie error", func(t *testing.T) {
		createdMovieReq := model.Movie{Title: "Test Film"}
		jsonStr, _ := json.Marshal(createdMovieReq)
		req, _ := http.NewRequest(http.MethodPost, "/moviecreate", bytes.NewBuffer(jsonStr))

		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)
		mh.CreateMovie(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	t.Run("Success", func(t *testing.T) {
		createdMovieReq := model.Movie{Title: "Test Film"}
		jsonStr, _ := json.Marshal(createdMovieReq)
		req, _ := http.NewRequest(http.MethodPost, "/moviecreate", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)
		mh.CreateMovie(rec, req)

		responseBodyStr := rec.Body.String()
		assert.Equal(t, "New movie is successfully added", responseBodyStr)
	})
}

func TestMovieHandler_DeleteMovie_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "GET not allowed",
			reqMethod: http.MethodGet,
		},
		{
			name:      "POST not allowed",
			reqMethod: http.MethodPost,
		},
		{
			name:      "PUT not allowed",
			reqMethod: http.MethodPut,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "/deletemovie", http.NoBody)
			rec := httptest.NewRecorder()

			mh := NewMovieHandler(nil)
			mh.DeleteMovie(rec, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
		})
	}
}

func TestMovieHandler_DeleteMovie(t *testing.T) {
	t.Run("delete movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/deletemovie", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)

		mh.DeleteMovie(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/deletemovie?id=1", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)

		mh.DeleteMovie(rec, req)

		assert.Equal(t, "Movie is successfully deleted", rec.Body.String())
	})
}

func TestMovieHandler_DeleteAllMovie_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "GET not allowed",
			reqMethod: http.MethodGet,
		},
		{
			name:      "POST not allowed",
			reqMethod: http.MethodPost,
		},
		{
			name:      "PUT not allowed",
			reqMethod: http.MethodPut,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "/deleteallmovies", http.NoBody)
			res := httptest.NewRecorder()

			mh := NewMovieHandler(nil)
			mh.DeleteAllMovies(res, req)

			assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
		})
	}
}

func TestMovieHandler_DeleteAllMovies(t *testing.T) {
	t.Run("delete all movie error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/deleteallmovies", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)
		mh.DeleteAllMovies(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
	t.Run("delete all movie successfully", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/deleteallmovies", http.NoBody)
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)
		mh.DeleteAllMovies(rec, req)

		assert.Equal(t, "All movies are successfully deleted", rec.Body.String())
	})
}

func TestMovieHandler_UpdateMovie_HTTPMethod_Tests(t *testing.T) {
	tests := []struct {
		name      string
		reqMethod string
	}{
		{
			name:      "GET not allowed",
			reqMethod: http.MethodGet,
		},
		{
			name:      "POST not allowed",
			reqMethod: http.MethodPost,
		},
		{
			name:      "DELETE not allowed",
			reqMethod: http.MethodDelete,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.reqMethod, "updatemovie", http.NoBody)
			rec := httptest.NewRecorder()

			mh := NewMovieHandler(nil)
			mh.UpdateMovie(rec, req)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)

		})
	}
}

func TestMovieHandler_UpdateMovie(t *testing.T) {
	t.Run("update movie error", func(t *testing.T) {
		updatedMovie := model.Movie{Title: "Test Movie"}
		jsonStr, _ := json.Marshal(updatedMovie)

		req, _ := http.NewRequest(http.MethodPatch, "/updatemovie", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: true}
		mh := NewMovieHandler(mockMovieSvc)
		mh.UpdateMovie(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("uptade movie successfully", func(t *testing.T) {
		updatedMovie := model.Movie{Title: "Test Movie"}
		jsonStr, _ := json.Marshal(updatedMovie)

		req, _ := http.NewRequest(http.MethodPatch, "/updatemovie", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		mockMovieSvc := service.MockMovieService{ReturnErr: false}
		mh := NewMovieHandler(mockMovieSvc)
		mh.UpdateMovie(rec, req)

		responseBodyStr := rec.Body.String()
		assert.Equal(t, "Movie is successfully updated", responseBodyStr)
	})
}
