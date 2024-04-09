package main

import (
	"fmt"
	"net/http"
	"log"

	"task-management-backend/github.com/anilmaurya61/internal/db"
	"task-management-backend/github.com/anilmaurya61/internal/routes"
)

func main(){

	db.ConnectionDb()
	r := routes.SetupRoutes() 
    fmt.Println("Server started at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", r))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    // Your handler logic here
	w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "Hello, World!")
}