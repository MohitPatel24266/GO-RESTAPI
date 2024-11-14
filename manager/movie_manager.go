// manager/movie_manager.go

package manager

import (
	"go_rest_mohit/request"
	"go_rest_mohit/response"
	"go_rest_mohit/services"
	"log"
)

// MovieManager handles business logic for movies
type MovieManager struct {
	movieService *services.MovieService
}

// NewMovieManager initializes a new MovieManager
func NewMovieManager(movieService *services.MovieService) *MovieManager {
	return &MovieManager{movieService: movieService}
}

// CreateMovie adds a new movie by interacting with the service
func (m *MovieManager) CreateMovie(movie *request.Request) (*response.Response, error) {
	log.Println("Mgr begin")

	// Call the service to create the movie
	res, err := m.movieService.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	log.Println("Mgr close")
	return res, nil
}

// GetAllMovies retrieves all movies by calling the service
// func (m *MovieManager) GetAllMovies() ([]response.Response, error) {
// 	return m.movieService.GetAllMovies()
// }

// func (m *MovieManager) GetAllMovies(limit, offset int) ([]response.Response, error) {
// 	// Here you can implement any additional business logic
// 	log.Println("Manager: GetAllMovies")
// 	return m.movieService.GetAllMovies(limit, offset)
// }

func (m *MovieManager) GetAllMovies(page, recordSize int) ([]response.Response, error) {
	// If page is -1, set recordSize to -1 to fetch all records
	if page == -1 {
		recordSize = -1
	}

	// Calculate offset if not fetching all records
	var offset int
	if recordSize != -1 {
		offset = (page - 1) * recordSize
	}

	// Call the MovieService to get the movies from the database
	return m.movieService.GetAllMovies(recordSize, offset)
}

// GetMovieByID retrieves a single movie by ID
func (m *MovieManager) GetMovieByID(id int) (*response.Response, error) {
	return m.movieService.GetMovieByID(id)
}

// UpdateMovie modifies an existing movie by calling the service
func (m *MovieManager) UpdateMovie(movie *request.Request) error {
	return m.movieService.UpdateMovie(movie)
}

// DeleteMovie removes a movie by ID by calling the service
func (m *MovieManager) DeleteMovie(id int) error {
	return m.movieService.DeleteMovie(id)
}
