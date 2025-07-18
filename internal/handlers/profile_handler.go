package handlers

import (
	"authenticator/internal/middleware"
	"authenticator/internal/services"
	"encoding/json"
	"net/http"
)

// the JSON structure for the user profile that too only email, not the password hash
type ProfileResponse struct {
	Email string `json:"email"`
}

// HTTP handler for the /getprofile endpoint, which assumes AuthMiddleware has already run and validated the user
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// to verify "GET"
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// user's email is passed via the request context by the middleware
	email, ok := middleware.GetUserEmailFromContext(r.Context())
	if !ok {
		// This should not happen if the middleware is working correctly
		http.Error(w, "Could not retrieve user email from context", http.StatusInternalServerError)
		return
	}

	// call to the service to get user details
	user, err := services.GetUserProfile(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	//response object creation
	res := ProfileResponse{
		Email: user.Email,
	}

	//Sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
