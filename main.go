package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/ReynekeJluc/go_study.git/db"
	routes "github.com/ReynekeJluc/go_study.git/routes"
)

func main() {
	database.ConnectDB()

	r := routes.SetupRouter()

	fmt.Println("starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}