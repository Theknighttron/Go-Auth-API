package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/polyhistor2050/Go-Auth-API/internal/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// Load the .env file
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Extract port number from dbURL
	dbParts := strings.Split(dbURL, ":")
	if len(dbParts) < 2 {
		log.Fatal("Invalid DB_URL format")
	}

	portParts := strings.Split(dbParts[1], "@")
	if len(portParts) < 2 {
		log.Fatal("Invalid DB_URL format")
	}

	fmt.Println("Port:", portParts[0]) // Debug statement
	fmt.Println("Port:", dbURL)

	// Make connection to the database
	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Can't connect to the database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// Create a new router object
	router := chi.NewRouter()

	// Cors configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/test", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	// Construct the server address
	addr := ":" + port

	// Create the server
	server := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	log.Printf("Server is running on port %v", dbURL)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
