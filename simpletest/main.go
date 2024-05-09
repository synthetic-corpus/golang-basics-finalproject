package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"strconv"
)

func helloAPI(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	reply := map[string]string{
		"Words": "Are over here",
		"These are": "more words",
		"Hello": "World"
	}

	json.NewEncoder(w).Encode(reply)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello",helloAPI).Methods("GET")

	port := ":" + strconv.Iota(os.Getenv("API_PORT"))
	fmt.Println("We are listening on " + port)
}