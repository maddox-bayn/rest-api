package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache map[int]string
var cacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)

	mux.HandleFunc("POST /user/{id}", CreateUser)
	mux.HandleFunc("GET /user/{id}", getUser)
	mux.HandleFunc("DELETE /user/{id}", deleteUser)
	fmt.Println("starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userCache[len(userCache)+1] = user.Name

}
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}
	cacheMutex.Lock()
	name, ok := userCache[id]
	cacheMutex.Unlock()
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(map[string]string{"name": name})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}
	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
