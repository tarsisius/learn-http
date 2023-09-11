package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in env")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	v1 := chi.NewRouter()
	v1.Get("/health", handleHealth)
	v1.Get("/error", handleError)

	router.Mount("/v1", v1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server running on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port", port)
}
