package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

var db = map[int64]Customer{
	1: {
		1,
		"Hans Meyer",
		"engineer",
		"hans.meyer@coolmail.com",
		"0171/123456789",
		true,
	},
	2: {
		2,
		"Jens Meyer",
		"service engineer",
		"jens.meyer@coolmail.com",
		"0171/987654321",
		false,
	},
	3: {
		3,
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
	if _, ok := db[customer.Id]; ok {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(emptyResponse)
	} else {
		db[customer.Id] = customer
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(db[customer.Id])
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
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
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
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
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
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

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	fmt.Println("Server starts on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
