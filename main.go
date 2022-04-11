package main

import (
	"github.com/dilaragorum/movie-go/handler"
	"github.com/dilaragorum/movie-go/repository"
	"github.com/dilaragorum/movie-go/service"
	"log"
	"net/http"
)

func main() {
	movieInMemoryRepository := repository.NewInMemoryMovieRepository()
	movieService := service.NewDefaultMovieService(movieInMemoryRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	http.HandleFunc("/movies", movieHandler.GetMovies)
	http.HandleFunc("/movie", movieHandler.GetMovie)
	http.HandleFunc("/createmovie", movieHandler.CreateMovie)
	http.HandleFunc("/deletemovie", movieHandler.DeleteMovie)
	http.HandleFunc("/deleteallmovies", movieHandler.DeleteAllMovies)
	http.HandleFunc("/updatemovie", movieHandler.UpdateMovie)

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
