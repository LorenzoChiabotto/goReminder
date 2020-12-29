package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	"./authentication"
	"./database"
)

func addReminder(w http.ResponseWriter, r *http.Request) {
	var request string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)
}
func editReminder(w http.ResponseWriter, r *http.Request) {
	var request string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)

}
func deleteReminder(w http.ResponseWriter, r *http.Request) {
	var request string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(request)

}
func getReminders(w http.ResponseWriter, r *http.Request) {
	var placeholder authentication.Token

	err := authentication.ValidateToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//err = json.NewDecoder(r.Body).Decode(&placeholder)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)
}

func getUserToken(w http.ResponseWriter, r *http.Request) {
	var request authentication.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := authentication.GenerateToken(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
func addUser(w http.ResponseWriter, r *http.Request) {
	var request User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.Register_user(request.Name, request.Password, request.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(token2)
}
func getUser(w http.ResponseWriter, r *http.Request) {}

func main() {
	err := godotenv.Load("./.env", "./database/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/validate", getUserToken).Methods("POST")

	r.HandleFunc("/users", getUser).Methods("GET")
	r.HandleFunc("/users", addUser).Methods("POST")

	r.HandleFunc("/reminders", addReminder).Methods("POST")
	r.HandleFunc("/reminders", editReminder).Methods("PUT")
	//	r.HandleFunc("/reminders", deleteReminder).Methods("POST")
	r.HandleFunc("/reminders", getReminders).Methods("GET")

	http.ListenAndServe(":80", r)
}
