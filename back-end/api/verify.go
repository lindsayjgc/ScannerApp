package main

import (
	"encoding/json"
	"net/http"
)

type Email struct {
	Email string `json:"email"`
}

func VerifyEmailSignup(w http.ResponseWriter, r *http.Request) {
	IssueCode(w, r, "signup")
}

func IssueCode(w http.ResponseWriter, r *http.Request, emailType string) {
	w.Header().Set("Content-Type", "application/json")

	// Decode email from frontend
	var email Email
	err = json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	randomCode := GenerateRandomCode()
	err = SendEmail(email.Email, randomCode, "signup")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Verification email sent successfully"))
}