package handlers

import (
	"authenticator/internal/services"
	"encoding/json"
	"net/http"
)

// the expected JSON structure for a registration request
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// the JSON structure for a successful registration response
type RegisterResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

// the HTTP handler for the /register endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// verifying if it is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decoding the JSON request body
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.RegisterUser(req.Email, req.Password)
	if err != nil {
		//checking the error type to give a more specific error message
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	//success response creation
	res := RegisterResponse{
		Message: "User registered successfully",
		Email:   user.Email,
	}

	// to set the content type header and then sending the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(res)
}
