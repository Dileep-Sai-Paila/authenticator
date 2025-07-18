package handlers

import (
	"authenticator/internal/services"
	"encoding/json"
	"net/http"
	"time"
)

// the expected JSON for a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// the JSON for a successful login
type LoginResponse struct {
	Message string `json:"message"`
}

// the HTTP handler for the /login endpoint
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//call the service to attempt to log the user in
	tokenString, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	//if login is successful then only it set the token as an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true, //to prevent the client-side scripts from accessing the cookie
	})

	// Sending the success response
	res := LoginResponse{Message: "Logged in successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
