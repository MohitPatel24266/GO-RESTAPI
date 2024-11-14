// package route

// import (
// 	"go_rest_mohit/controller"

// 	"database/sql"

// 	"github.com/labstack/echo/v4"
// )

// func RegisterMovieRoutes(e *echo.Echo, db *sql.DB) {
// 	mc := &controller.MovieController{DB: db}

// 	e.POST("/movies/create", mc.CreateMovie)
// 	e.GET("/movies/:id", mc.GetMovie)
// 	e.GET("/movies", mc.GetAllMovies)
// 	e.PUT("/movies/:id", mc.UpdateMovie)
// 	e.DELETE("/movies/:id", mc.DeleteMovie)
// }


package route

import (
	"go_rest_mohit/controller"  // Correct import for controller package
	"github.com/labstack/echo/v4"
)

// SetupRoutes defines all the routes for the movie API
func SetupRoutes(e *echo.Echo) {
	e.POST("/movies/create", controller.CreateMovie)
	e.GET("/movies", controller.GetAllMovies)
	e.GET("/movies/:id", controller.GetMovie)
	e.PUT("/movies/:id", controller.UpdateMovie)
	e.DELETE("/movies/:id", controller.DeleteMovie)
}
