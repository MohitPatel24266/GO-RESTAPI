// services/movie_service.go

package services

import (
	"database/sql"
	"go_rest_mohit/request"
	"go_rest_mohit/response"
	"log"
)

// var res  *response.Response

// func res_ctrl(r *response.Response) {
// 	res = r
// }

// MovieService provides database-related operations for movies
type MovieService struct {
	DB *sql.DB
}

// NewMovieService initializes a new MovieService
func NewMovieService(db *sql.DB) *MovieService {
	return &MovieService{DB: db}
}

// CreateMovie adds a new movie to the database
func (s *MovieService) CreateMovie(movie *request.Request) (*response.Response, error) {
	log.Println("Service begin")

	var id int
	err := s.DB.QueryRow("INSERT INTO movies (name, genre, price) VALUES ($1, $2, $3) RETURNING id",
		movie.Name, movie.Genre, movie.Price).Scan(&id)
	if err != nil {
		return nil, err
	}

	res := &response.Response{
		ID:    id,
		Name:  movie.Name,
		Genre: movie.Genre,
		Price: movie.Price,
	}

	log.Println("Service close")
	return res, nil
}

func (s *MovieService) GetAllMovies(recordSize, offset int) ([]response.Response, error) {
	log.Println("Service: GetAllMovies")

	// If recordSize is -1, fetch all records
	if recordSize == -1 {
		rows, err := s.DB.Query("SELECT id, name, genre, price FROM movies")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var movies []response.Response
		for rows.Next() {
			var movie response.Response
			if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
				return nil, err
			}
			movies = append(movies, movie)
		}
		return movies, nil
	}

	// Query the database with pagination (LIMIT and OFFSET)
	rows, err := s.DB.Query("SELECT id, name, genre, price FROM movies LIMIT $1 OFFSET $2", recordSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []response.Response
	for rows.Next() {
		var movie response.Response
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
// GetMovieByID retrieves a single movie by ID
func (s *MovieService) GetMovieByID(id int) (*response.Response, error) {
	var movie response.Response
	err := s.DB.QueryRow("SELECT id, name, genre, price FROM movies WHERE id = $1", id).
		Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// UpdateMovie modifies an existing movie
func (s *MovieService) UpdateMovie(movie *request.Request) error {
	_, err := s.DB.Exec("UPDATE movies SET name=$1, genre=$2, price=$3 WHERE id=$4",
		movie.Name, movie.Genre, movie.Price, movie.ID)
	return err
}

// DeleteMovie removes a movie by ID
func (s *MovieService) DeleteMovie(id int) error {
	_, err := s.DB.Exec("DELETE FROM movies WHERE id=$1", id)
	return err
}
