package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

 	jwtRequest "github.com/dgrijalva/jwt-go/request"

	"./common"

	"./Reminders"
	)

func addReminder(w http.ResponseWriter, r *http.Request) {
	var request AddReminderRequest
	tokenString, err := jwtRequest.HeaderExtractor{"Authorization"}.ExtractToken(r)
	if err != nil {
		fmt.Println(err)
	}

	token, err := commonUtils.ValidateToken(tokenString)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = Reminders.AddReminder(Reminders.Reminder{
		Day:              request.Day,
		Repeat:           request.Repeat,
		Message:          request.Message,
	},
	token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
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
	err = json.NewEncoder(w).Encode(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

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
	err = json.NewEncoder(w).Encode(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func getReminders(w http.ResponseWriter, r *http.Request) {
	var request string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}}

func main() {
	err := godotenv.Load("./API/database/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/reminders", addReminder).Methods("POST")
	r.HandleFunc("/reminders", editReminder).Methods("PUT")
	r.HandleFunc("/reminders", deleteReminder).Methods("DELETE")
	r.HandleFunc("/reminders", getReminders).Methods("GET")

	log.Println("listening port 85")
	err = http.ListenAndServe(":85", r)

	if err != nil {
		log.Fatal("There was an error when serving the server", err)
	}
}
