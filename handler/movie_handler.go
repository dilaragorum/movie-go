package handler

import (
	"encoding/json"
	"errors"
	"github.com/dilaragorum/movie-go/model"
	"github.com/dilaragorum/movie-go/service"
	"net/http"
	"strconv"
)

// Önemli : Biz burada Handler struct'ı içerisinde getMovies / getMovie gibi metotları
// service.IMovieService'ı satisfy etmek için yazmıyoruz. Bu Dependency Injection oluyor.
// Ben Dependency Injection yaparak, IMovieService interface'ini satisfy eden struct'ların
//metotlarını çağırabiliyorum.
type movieHandler struct {
	service service.IMovieService
}

// Pointer kullanmak yerine bu metodu kullanıp içeride logic/ yada başka bir metot
// kullanabiliyorum. Böylece testlerde, newMovieHandler kullanıldığı için, bu kodum hepsinde
// çalıştırılıyor. MockMovie Struct'ı da DefaultMovieService struct'ıda olsa interface'i
// satisfy ettiği sürece kullanılabilir.
func NewMovieHandler(ms service.IMovieService) *movieHandler {
	return &movieHandler{service: ms}
}

// curl localhost:8080/movies | jq
func (mh *movieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// işi Service'a delege ediyoruz.
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

// curl "localhost:8080/movie?id=1" | jq
func (mh *movieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	// İşi service'a delege ediyoruz.
	movie, err := mh.service.GetMovie(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

/*
curl -X POST localhost:8080/moviecreate \
-H 'Content-Type: application/json' \
-d '{ "title": "Güzel film" }'
*/
func (mh *movieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "error when decoding json", http.StatusInternalServerError)
		return
	}

	// İşi Service'e delege ediyoruz.
	msg, err := mh.service.CreateMovie(movie)
	if err != nil {
		if errors.Is(err, service.ErrTitleIsNotEmpty) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Response Body'e yazıyor. byte olarak yazdırıyorum Body'e.
	w.Write([]byte(msg))
}

func (mh *movieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	// İşi service'e delege ediyorum.
	message, err := mh.service.DeleteMovie(id)

	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response body'e byte olarak yazdırdım.
	w.Write([]byte(message))
}

func (mh *movieHandler) DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//Service'a delegate ediyorum
	msg, err := mh.service.DeleteAllMovie()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(msg))
}

//curl -X PATCH "localhost:8080/updatemovie?id=1" -d '{ "title": "Güzel film" }'
func (mh *movieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "error when decoding json", http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	// id ve requestteki  body'i decode ederek işi Service' a delege ediyorum.
	msg, err := mh.service.UpdateMovie(id, movie)

	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) || errors.Is(err, service.ErrTitleIsNotEmpty) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(msg))
}
