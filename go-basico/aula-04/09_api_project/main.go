package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/diego/go-basico-01/handler"
	"github.com/diego/go-basico-01/middleware"
)

func main() {
	r := mux.NewRouter()

	// Aplicando o middleware de logging
	loggedMux := middleware.LoggingMiddleware(r)

	// Definindo as rotas e associando os handlers
	r.HandleFunc("/users/{id}", handler.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/users", handler.ListUsers).Methods(http.MethodGet)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
