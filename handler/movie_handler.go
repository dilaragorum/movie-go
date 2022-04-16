package handler

import (
	"encoding/json"
	"errors"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type movieHandler struct {
	service service.IMovieService
}

func NewMovieHandler(ms service.IMovieService) *movieHandler {
	return &movieHandler{service: ms}
}

// curl localhost:8080/movies | jq
func (mh *movieHandler) GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := mh.service.GetMovies()
	if err != nil {
		http.Error(w, "Unable to get all movies", http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

// curl "localhost:8080/movies/1" | jq
func (mh *movieHandler) GetMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	movie, err := mh.service.GetMovie(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrMovieNotFound) { // Test yaz
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

/*
curl -X POST localhost:8080/movies \
-H 'Content-Type: application/json' \
-d '{ "title": "A New Movie" }'
*/
func (mh *movieHandler) CreateMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "error when decoding json", http.StatusInternalServerError)
		return
	}

	err = mh.service.CreateMovie(movie)
	if err != nil {
		if errors.Is(err, service.ErrTitleIsNotEmpty) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Movie is successfully created"))
}

func (mh *movieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	err := mh.service.DeleteMovie(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) || errors.Is(err, service.ErrMovieNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Movie has been deleted succesfully"))
}

func (mh *movieHandler) DeleteAllMovies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := mh.service.DeleteAllMovie()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Movies successfully deleted"))
}

//curl -X PATCH "localhost:8080/movies/1" -d '{ "title": "Beautiful film" }'
func (mh *movieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "error when decoding json", http.StatusInternalServerError)
		return
	}

	err = mh.service.UpdateMovie(id, movie)

	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) ||
			errors.Is(err, service.ErrTitleIsNotEmpty) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrMovieNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Updated"))
}
