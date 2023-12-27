package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = map[int]string{}

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

func getJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []Job

	jobs = append(jobs, Job{ID: 1, Name: "Software Engineer"})
	jobs = append(jobs, Job{ID: 2, Name: "Seiyuu"})
	jobs = append(jobs, Job{ID: 3, Name: "Babysitter"})
	jobs = append(jobs, Job{ID: 4, Name: "AI Engineer"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
