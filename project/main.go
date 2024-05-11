package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	//"net/mail"
)

// Structs section
type Customer struct {
	ID string // always a google uuid
	Name string
	Role string
	Email string
	Phone string
	Contacted bool
}

// The fake Database
type Database struct {
	Customers map[string]Customer
}

// The Database Level CRRUD methods

func (db *Database) writeCustomer(newCustomer Customer) bool {
	db.Customers[newCustomer.ID] = newCustomer
	return true // will put this into a try catch block later
}

func (db *Database) retrieveAll() []Customer {
	returnThis := []Customer{}
	for key := range db.Customers{
		returnThis = append(returnThis, db.Customers[key])
	}
	return returnThis
}

func (db *Database) retrieveOne(ID string) Customer {
	return db.Customers[ID] // will add better error handling later
}

func(db *Database) updateOne(ID string, customer Customer){
	// Logic used here should prevent users from using this to alter the UUID.
	customer.ID = ID
	db.Customers[ID] = customer
}

func(db *Database) deleteOne(ID string){
	delete(db.Customers, ID)
}

// Instatiates an empty Database
var myFakeDatabase Database = Database{
	Customers: map[string]Customer{},
}


// The API fucntions
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
	
	user_id := mux.Vars(r)["id"]
	if _, ok := myFakeDatabase.Customers[user_id]; ok{
		w.WriteHeader(http.StatusOK)
		reply := myFakeDatabase.retrieveOne(user_id)
		json.NewEncoder(w).Encode(reply)
	}else{
		w.WriteHeader(http.StatusNotFound)
		reply := map[string]string{
			"Message": "User ID was not found!",
			"ID": user_id,
		}
		json.NewEncoder(w).Encode(reply)
	}
}

// retieve users
func retrieveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := myFakeDatabase.retrieveAll()

	json.NewEncoder(w).Encode(reply)
}

// update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user_id := mux.Vars(r)["id"]

	if _, ok := myFakeDatabase.Customers[user_id]; ok{
		w.WriteHeader(http.StatusOK)
		reply := map[string]string{
			"Message":"Update path found a user!",
			"ID": user_id,
		}
		json.NewEncoder(w).Encode(reply)
	}else{
		w.WriteHeader(http.StatusNotFound)
		reply := map[string]string{
			"Message": "User ID was not found!",
			"ID": user_id,
		}
		json.NewEncoder(w).Encode(reply)
	}
}

// delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	user_id := mux.Vars(r)["id"]
	if _, ok := myFakeDatabase.Customers[user_id]; ok{
		w.WriteHeader(http.StatusOK)
		reply := map[string]string{
			"Message":"Delete path found a user!",
			"ID": user_id,
		}
		json.NewEncoder(w).Encode(reply)
	}else{
		w.WriteHeader(http.StatusNotFound)
		reply := map[string]string{
			"Message": "User ID was not found!",
			"ID": user_id,
		}
		json.NewEncoder(w).Encode(reply)
	}
}

func main() {
	

// Populates the Database
	var myFakeUsers []Customer = []Customer{
		{	
			ID: uuid.NewString(),
			Name: "John Doe",
			Role: "Buyer",
			Email: "yourEmail@google.net",
			Phone: "818-555-1515",
			Contacted: false,
		},
		{
			ID: uuid.NewString(),
			Name: "Allison Looper",
			Role: "Test Merchant",
			Email: "allyloop1990@gmail.com",
			Phone: "626-555-8055",
			Contacted: false,
		},
		{
			ID: uuid.NewString(),
			Name: "Holly Friedman",
			Role: "Another person",
			Email: "Hollyofthevalley@protonmail.com",
			Phone: "818-555-7112",
			Contacted: false,
		},
	}
	// Populate database
	for _,person := range myFakeUsers{
		myFakeDatabase.writeCustomer(person)
	}
	fmt.Println(myFakeDatabase.Customers)

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
