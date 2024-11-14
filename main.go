package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"go_rest_mohit/controller"
	"go_rest_mohit/manager"
	"go_rest_mohit/route"
	"go_rest_mohit/services"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Database connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Service, Manager, and Controller
	movieService := services.NewMovieService(db)         // Service handles database operations
	movieManager := manager.NewMovieManager(movieService) // Manager handles business logic
	controller.InitializeController(movieManager)         // Controller handles HTTP requests

	// Create Echo instance and setup routes
	e := echo.New()
	route.SetupRoutes(e) // Setup routes using the route package

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
