package routes 

import (
    "net/http"

    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", handlerFunc).Methods("GET")
 
	
    return r
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    // Your handler logic here
}
