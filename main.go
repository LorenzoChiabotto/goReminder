package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	"./authentication"
)

func addReminder(w http.ResponseWriter, r *http.Request) {
	var placeholder string
	err := json.NewDecoder(r.Body).Decode(&placeholder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)
}
func editReminder(w http.ResponseWriter, r *http.Request) {
	var placeholder string
	err := json.NewDecoder(r.Body).Decode(&placeholder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)

}
func deleteReminder(w http.ResponseWriter, r *http.Request) {
	var placeholder string
	err := json.NewDecoder(r.Body).Decode(&placeholder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)

}
func getReminders(w http.ResponseWriter, r *http.Request) {
	var placeholder string
	err := json.NewDecoder(r.Body).Decode(&placeholder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	token, err := authentication.GenerateToken()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
func getUserToken(w http.ResponseWriter, r *http.Request) {
	var token string
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token2, err := authentication.ValidateToken(token)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token2)
}
func getUser(w http.ResponseWriter, r *http.Request) {}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/validate", getUserToken).Methods("POST")

	r.HandleFunc("/users", getUser).Methods("GET")
	r.HandleFunc("/users", addUser).Methods("POST")

	r.HandleFunc("/{userName}/reminders", addReminder).Methods("POST")
	r.HandleFunc("/{userName}/reminders", editReminder).Methods("PUT")
	r.HandleFunc("/{userName}/reminders", deleteReminder).Methods("POST")
	r.HandleFunc("/{userName}/reminders", getReminders).Methods("GET")

	http.ListenAndServe(":80", r)
}
