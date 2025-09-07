package routes

import (
	"fmt"
	"net/http"

	handlers "github.com/ReynekeJluc/go_study.git/handlers"
	mux "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Working!")
	})

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	api.HandleFunc("/books/{BookId}", handlers.GetBook).Methods("GET")
	api.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	api.HandleFunc("/books/{BookId}", handlers.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{BookId}", handlers.DeleteBook).Methods("DELETE")

	return r
}
