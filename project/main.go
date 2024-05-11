package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func helloAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reply := map[string]string{
		"Words":     "Are over here",
		"These are": "more words",
		"Hello":     "World",
	}

	json.NewEncoder(w).Encode(reply)
}

// create user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	reply := map[string]string{
		"Message": "Create User Called as expected",
	}

	json.NewEncoder(w).Encode(reply)
}

// retrieve user
func retrieveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := map[string]string{
		"Message": "Retrieve User Called as expected",
	}

	json.NewEncoder(w).Encode(reply)
}

// retieve users
func retrieveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := map[string]string{
		"Message": "Retrieve ALL Users Called as expected",
	}

	json.NewEncoder(w).Encode(reply)
}

// update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := map[string]string{
		"Message": "Update User Called as expected",
	}

	json.NewEncoder(w).Encode(reply)
}

// delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := map[string]string{
		"Message": "Delete user Called as expected",
	}

	json.NewEncoder(w).Encode(reply)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloAPI).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", retrieveUser).Methods("GET")
	router.HandleFunc("/users", retrieveUsers).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	port := ":" + os.Getenv("API_PORT")
	fmt.Println("We are listening on " + port)

	http.ListenAndServe(port, router)
}
