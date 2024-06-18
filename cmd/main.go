package main

import (
	"go-poc-example/internal/api"
	"go-poc-example/internal/repository/memdb"
	"go-poc-example/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the repository
	memdb := repository.NewMemDB()

	// Initialize the service with the repository
	svc := service.NewService(memdb)

	// Initialize the handler with the service
	handler := api.NewHandler(svc)

	// Create a new ServeMux
	router := mux.NewRouter()

	// Register routes
	handler.RegisterRoutes(router)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
