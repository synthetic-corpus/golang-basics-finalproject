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
	Phone uint64
	Contacted bool
}

type CustomerUpdate struct {
	Name string
	Role string
	Email string
	Phone uint64
	Contact bool
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

func(db *Database) updateOne(ID string, update Customer){
	// Api function already has santized the input of ID, so we will find a user to update
	// Doesn't update by field, so fields can be set to null. Better updates could be had in future versions.
	update.ID = ID // If a bad reqeust tried to over write the UUID, than this will correct it.
	db.Customers[ID] = update
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
func addCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newCustomer Customer
	err := decoder.Decode(&newCustomer)
	if err != nil { panic(err)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	newCustomer.ID = uuid.NewString()
	myFakeDatabase.writeCustomer(newCustomer)
	json.NewEncoder(w).Encode(newCustomer)
}

// retrieve user
func getCustomer(w http.ResponseWriter, r *http.Request) {
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
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	reply := myFakeDatabase.retrieveAll()

	json.NewEncoder(w).Encode(reply)
}

// update user
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user_id := mux.Vars(r)["id"]

	if _, ok := myFakeDatabase.Customers[user_id]; ok{
		
		decoder := json.NewDecoder(r.Body)
		var updateCustomer Customer
		err := decoder.Decode(&updateCustomer)
		if err != nil { panic(err)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		myFakeDatabase.updateOne(user_id,updateCustomer)
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
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	user_id := mux.Vars(r)["id"]
	if _, ok := myFakeDatabase.Customers[user_id]; ok{
		w.WriteHeader(http.StatusOK)
		myFakeDatabase.deleteOne(user_id)
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

// Home Page Function
func showHomePage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "./homepage.html")
}

func main() {
	

// Populates the Database
	var myFakeUsers []Customer = []Customer{
		{	
			ID: uuid.NewString(),
			Name: "John Doe",
			Role: "Buyer",
			Email: "yourEmail@google.net",
			Phone: 8185551515,
			Contacted: false,
		},
		{
			ID: uuid.NewString(),
			Name: "Allison Looper",
			Role: "Test Merchant",
			Email: "allyloop1990@gmail.com",
			Phone: 6265558055,
			Contacted: false,
		},
		{
			ID: uuid.NewString(),
			Name: "Holly Friedman",
			Role: "Another person",
			Email: "Hollyofthevalley@protonmail.com",
			Phone: 818555112,
			Contacted: false,
		},
	}
	// Populate database
	for _,person := range myFakeUsers{
		myFakeDatabase.writeCustomer(person)
	}
	fmt.Println(myFakeDatabase.Customers)

	

	router := mux.NewRouter()
	router.HandleFunc("/", showHomePage).Methods("GET")
	router.HandleFunc("/hello", helloAPI).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	port := ":" + os.Getenv("API_PORT")
	fmt.Println("We are listening on " + port)

	http.ListenAndServe(port, router)
}
