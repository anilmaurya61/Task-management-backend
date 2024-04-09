package main

import (
	"fmt"
	"net/http"
	"log"

	"task-management-backend/github.com/anilmaurya61/internal/db"

	"github.com/gorilla/mux"
)

func main(){

	db.ConnectionDb()

    r := mux.NewRouter()

    r.HandleFunc("/", handlerFunc).Methods("GET")

    fmt.Println("Server started at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", r))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    // Your handler logic here
	w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "Hello, World!")
}