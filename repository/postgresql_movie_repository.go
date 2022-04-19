package repository

import (
	"database/sql"
	"fmt"
	"github.com/dilaragorum/movie-go/model"
	_ "github.com/lib/pq"
	"log"
)

type postgresqlMovieRepository struct {
	connectionPool *sql.DB
}

func NewPostgreSQLMovieRepository() *postgresqlMovieRepository {
	// Template: "postgresql://<username>:<password>@<database_ip>/<database-name>?sslmode=disable
	connStr := "postgres://postgres:postgres@localhost/movie-db?sslmode=disable"
	connectionPool, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = connectionPool.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL connection is successful")

	return &postgresqlMovieRepository{
		connectionPool: connectionPool,
	}
}

func (p *postgresqlMovieRepository) GetMovies() ([]model.Movie, error) {
	rows, err := p.connectionPool.Query("SELECT * FROM movies")
	if err != nil {
		return []model.Movie{}, err
	}
	defer rows.Close()

	movies := make([]model.Movie, 0)

	for rows.Next() {
		mv := model.Movie{}
		err := rows.Scan(&mv.ID, &mv.Title, &mv.ReleaseYear, &mv.Score)
		if err != nil {
			return movies, err
		}
		movies = append(movies, mv)
	}

	return movies, nil
}

func (p *postgresqlMovieRepository) GetMovie(id int) (model.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresqlMovieRepository) CreateMovie(movie model.Movie) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgresqlMovieRepository) DeleteMovie(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgresqlMovieRepository) DeleteAllMovies() error {
	//TODO implement me
	panic("implement me")
}

func (p *postgresqlMovieRepository) UpdateMovie(id int, movie model.Movie) error {
	//TODO implement me
	panic("implement me")
}
