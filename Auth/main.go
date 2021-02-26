package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"./authentication"
)

func main() {
	err := godotenv.Load("./database/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/auth", getToken).Methods("POST")

	http.ListenAndServe(":80", r)
}

func getToken(w http.ResponseWriter, r *http.Request) {
	var request RequestToken
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := authentication.GenerateToken(request.User)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(
		ResponseToken{
			Token:*token,
		})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}