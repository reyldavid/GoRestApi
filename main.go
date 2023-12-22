package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = map[int]string{}

func main() {
	fmt.Println("Hello, Alison!")

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", createUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", readUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", updateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8088", r))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data map[string]string

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users[id] = data["name"]

	w.WriteHeader(http.StatusCreated)
}

func readUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name, ok := users[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := map[string]string{"name": name}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data map[string]string

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users[id] = data["name"]
	w.WriteHeader(http.StatusOK)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	delete(users, id)
	w.WriteHeader(http.StatusOK)
}
