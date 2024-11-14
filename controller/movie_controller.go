// controller/movie_controller.go

package controller

import (
	"database/sql"
	// "go_rest_mohit/response"
	// "go_rest_mohit/services"
	"github.com/labstack/echo/v4"
	"go_rest_mohit/manager"
	"go_rest_mohit/request"
	"log"
	"net/http"
	"strconv"
)

// // Declare the MovieManager
// var movieManager *services.MovieService

// // InitializeController initializes the controller with a MovieService
// func InitializeController(service *services.MovieService) {
// 	movieManager = service
// }

var movieManager *manager.MovieManager
// var movieService *services.MovieService

func InitializeController(mgr *manager.MovieManager) {
	movieManager = mgr
}

// func InitializeControllerService(service *services.MovieService) {
// 	movieService = service
// }

// CreateMovie handles the creation of a new movie
func CreateMovie(c echo.Context) error {
	log.Println("ctr begin")

	req := new(request.Request)

	// Bind the request data to the Movie struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate the movie data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call the MovieManager to create the movie
	r, err := movieManager.CreateMovie(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Println("ctr close")

	// Return the response directly from the service/manager
	return c.JSON(http.StatusCreated, r)
}

func GetAllMovies(c echo.Context) error {
	//Converting QueryParams String to Interger
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1 // Default to page 1 if invalid
	}

	recordSize, err := strconv.Atoi(c.QueryParam("recordsize"))
	if err != nil || recordSize <= 0 {
		recordSize = recordSize*1 
	}

	if page == -1 {
		recordSize = -1 // Indicating to fetch all records
	}

	// Call the Manager to fetch paginated movies
	movies, err := movieManager.GetAllMovies(page, recordSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

// GetMovie fetches a movie by ID
func GetMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	movie, err := movieManager.GetMovieByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Movie not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// res := response.Response{
	// 	ID:    movie.ID,
	// 	Name:  movie.Name,
	// 	Genre: movie.Genre,
	// 	Price: movie.Price,
	// }
	return c.JSON(http.StatusOK, movie)
}

// UpdateMovie handles updating an existing movie
func UpdateMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	req := new(request.Request)

	// Bind the request data to the Movie struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	req.ID = id

	// Validate the movie data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call the MovieManager to update the movie
	err = movieManager.UpdateMovie(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, req)
}

// DeleteMovie deletes a movie by ID
func DeleteMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	err = movieManager.DeleteMovie(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Movie deleted successfully"})
}
