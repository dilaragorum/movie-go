package model

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Revenue     float64 `json:"revenue"`
	Budget      float64 `json:"budget"`
	ReleaseYear int     `json:"release_year"`
	Score       float64 `json:"score"`
}
