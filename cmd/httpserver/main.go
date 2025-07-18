package main

import (
	"authenticator/internal/handlers"
	"authenticator/internal/middleware"
	"authenticator/internal/repositories"
	"log"
	"net/http"
)

func main() {
	//database connection through my account in "MongoDB Atlas"
	const mongoURI = "mongodb+srv://dileepsaipaila:ThnhdLx486kjbuXw@cluster0.boe2nyk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	repositories.ConnectDB(mongoURI)

	//server setup
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running!"))
	})

	//public routes
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	//protected route
	profileHandler := http.HandlerFunc(handlers.GetProfileHandler)
	//now i am wrapping this handler with the AuthMiddleware function which is in my auth_middleware.go inside authenticator/internal/middleware.

	mux.Handle("/getprofile", middleware.AuthMiddleware(profileHandler))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
