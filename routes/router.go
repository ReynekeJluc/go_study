package routes

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/ReynekeJluc/go_study.git/db"
	handlers "github.com/ReynekeJluc/go_study.git/handlers"
	middleware "github.com/ReynekeJluc/go_study.git/middlewares"
	mux "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
		defer database.DB.Close()
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Working!")
	})

	auth_router := r.PathPrefix("/api/auth").Subrouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)


	auth_router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	auth_router.HandleFunc("/refresh", handlers.RefreshAccessTokenHandler).Methods("POST")
	auth_router.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	api.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	api.HandleFunc("/books/{BookId}", handlers.GetBook).Methods("GET")
	api.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	api.HandleFunc("/books/{BookId}", handlers.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{BookId}", handlers.DeleteBook).Methods("DELETE")

	return r
}
