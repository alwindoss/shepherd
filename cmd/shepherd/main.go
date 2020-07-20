package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type tokenResponseObj struct {
	Token string `json:"token,omitempty"`
}

// User is the object that holds user data
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type userResponseObj struct {
	User *User `json:"user,omitempty"`
}

func main() {
	router := mux.NewRouter()
	// apiRoute := router.PathPrefix("/api")
	// cors := handlers.CORS(
	// allowedHeaders := handlers.AllowedHeaders([]string{"*"})
	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedCredentials := handlers.AllowCredentials()
	// )
	// headersOK := handlers.AllowedHeaders([]string{"*"})
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	credsOK := handlers.AllowCredentials()
	// router.Use(cors)
	router.Path("/api/auth/login").Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := &tokenResponseObj{
			Token: "sdafasdfasdfasdfdf",
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Unknown error", http.StatusInternalServerError)
			return
		}
		return
	})
	router.Path("/api/auth/login").Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := &tokenResponseObj{
			Token: "sdafasdfasdfasdfdf",
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Unknown error", http.StatusInternalServerError)
			return
		}
		return
	})
	router.Path("/api/auth/user").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		usr := &User{
			Name:  "Alwin Doss",
			Email: "alwindoss84@gmail.com",
		}
		resp := &userResponseObj{
			User: usr,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Unknown error", http.StatusInternalServerError)
			return
		}
		return

	})
	// router.
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOK, originsOK, methodsOK, credsOK)(router)))
}
