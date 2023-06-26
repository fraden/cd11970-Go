package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/citylist", cityList)

	http.ListenAndServe(":3000", nil)

}
