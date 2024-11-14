package model

import (
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Movie struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" validate:"required,min=1,max=100"`
	Genre string  `json:"genre" validate:"required,min=1,max=50"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

// Validate the Movie struct
func (m *Movie) Validate() error {
	return validate.Struct(m)
}

// Create a new movie in the database
func CreateMovie(db *sql.DB, movie *Movie) error {
	if err := movie.Validate(); err != nil {
		return err
	}
	return db.QueryRow("INSERT INTO movies (name, genre, price) VALUES ($1, $2, $3) RETURNING id",
		movie.Name, movie.Genre, movie.Price).Scan(&movie.ID)
}

// Get a movie by ID
func GetMovieByID(db *sql.DB, id int) (*Movie, error) {
	movie := &Movie{}
	err := db.QueryRow("SELECT id, name, genre, price FROM movies WHERE id = $1", id).
		Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price)
	if err == sql.ErrNoRows {
		return nil, errors.New("movie not found")
	}
	return movie, err
}

// Get all movies from the database
func GetAllMovies(db *sql.DB) ([]Movie, error) {
	rows, err := db.Query("SELECT id, name, genre, price FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// Update a movie's details
func UpdateMovie(db *sql.DB, movie *Movie) error {
	if err := movie.Validate(); err != nil {
		return err
	}
	_, err := db.Exec("UPDATE movies SET name=$1, genre=$2, price=$3 WHERE id=$4",
		movie.Name, movie.Genre, movie.Price, movie.ID)
	return err
}

// Delete a movie by ID
func DeleteMovie(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM movies WHERE id=$1", id)
	return err
}


// package model

// import (
// 	"database/sql"
// 	"errors"
// 	// "github.com/go-playground/validator/v10"
// )

// type Movie struct {
// 	ID    int     `json:"id"`
// 	Name  string  `json:"name"`
// 	Genre string  `json:"genre"`
// 	Price float64 `json:"price"`
// }

// // Create a new movie in the database
// func CreateMovie(db *sql.DB, movie *Movie) error {
// 	return db.QueryRow("INSERT INTO movies (name, genre, price) VALUES ($1, $2, $3) RETURNING id",
// 		movie.Name, movie.Genre, movie.Price).Scan(&movie.ID)
// }

// // Get a movie by ID
// func GetMovieByID(db *sql.DB, id int) (*Movie, error) {
// 	movie := &Movie{}
// 	err := db.QueryRow("SELECT id, name, genre, price FROM movies WHERE id = $1", id).
// 		Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price)
// 	if err == sql.ErrNoRows {
// 		return nil, errors.New("movie not found")
// 	}
// 	return movie, err
// }

// // Get all movies from the database
// func GetAllMovies(db *sql.DB) ([]Movie, error) {
// 	rows, err := db.Query("SELECT id, name, genre, price FROM movies")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var movies []Movie
// 	for rows.Next() {
// 		var movie Movie
// 		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
// 			return nil, err
// 		}
// 		movies = append(movies, movie)
// 	}
// 	return movies, nil
// }

// // Update a movie's details
// func UpdateMovie(db *sql.DB, movie *Movie) error {
// 	_, err := db.Exec("UPDATE movies SET name=$1, genre=$2, price=$3 WHERE id=$4",
// 		movie.Name, movie.Genre, movie.Price, movie.ID)
// 	return err
// }

// // Delete a movie by ID
// func DeleteMovie(db *sql.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM movies WHERE id=$1", id)
// 	return err
// }
