package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/tarsisius/learn-http/db"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *db.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect database: ", err)
	}

	apiConfig := ApiConfig{
		DB: db.New(conn),
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	r.Get("/server-ok", handleOk)
	r.Get("/server-error", handleError)

	r.Post("/create-user", apiConfig.handleCreateUser)
	r.Get("/get-user", apiConfig.middlewareAuth(apiConfig.handleGetUser))

	r.Post("/create-feed", apiConfig.middlewareAuth(apiConfig.handleCreateFeed))
	r.Get("/get-feeds", apiConfig.handleGetFeeds)

	r.Post("/follow", apiConfig.middlewareAuth((apiConfig.handleFollow)))
	r.Get("/get-follows", apiConfig.middlewareAuth(apiConfig.handleGetFollows))
	r.Delete("/delete-follow/{ID}", apiConfig.middlewareAuth(apiConfig.handleDeleteFollow))

	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	log.Printf("Server running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port", port)
}
