package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var db = map[string]Customer{
	"48364386-95be-47f1-9b3b-7c152469be0b": {
		"48364386-95be-47f1-9b3b-7c152469be0b",
		"Hans Meyer",
		"engineer",
		"hans.meyer@coolmail.com",
		"0171/123456789",
		true,
	},
	"d763da6f-1b79-410d-a243-35a1e2355995": {
		"d763da6f-1b79-410d-a243-35a1e2355995",
		"Jens Meyer",
		"service engineer",
		"jens.meyer@coolmail.com",
		"0171/987654321",
		false,
	},
	"39f9a196-1612-458a-9ab0-d64288b7f35b": {
		"39f9a196-1612-458a-9ab0-d64288b7f35b",
		"Daniel Meyer",
		"automation engineer",
		"daniel.meyer@coolmail.com",
		"0171/000000000",
		true,
	},
}

var emptyResponse = map[string]string{}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var customer Customer
	json.Unmarshal(body, &customer)
	id := uuid.New().String()
	customer.Id = id
	db[id] = customer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db[customer.Id])

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	if _, ok := db[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyResponse)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		var customer Customer
		json.Unmarshal(body, &customer)
		db[id] = customer
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(db[customer.Id])
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	if _, ok := db[id]; ok {
		delete(db, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyResponse)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]
	if _, ok := db[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(emptyResponse)
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func main() {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/", fileServer)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	fmt.Println("Server starts on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// writes a JSON response to the ResponseWriter
func writeJsonResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
