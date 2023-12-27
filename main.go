package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, Alison!")

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", createUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", readUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", updateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/jobs", getJobs).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
