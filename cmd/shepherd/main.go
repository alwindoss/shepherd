package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type tokenResponseObj struct {
	Token string `json:"token,omitempty"`
}

// User is the object that holds user data
type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type userResponseObj struct {
	User *User `json:"user,omitempty"`
}

var jwtKey = []byte("my_secret_key")

// Claims holds the claims
type Claims struct {
	User *User `json:"user"`
	jwt.StandardClaims
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalf("unable to create db")
	}
	router := mux.NewRouter()
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	credsOK := handlers.AllowCredentials()
	// router.Use(cors)
	router.Path("/api/auth/login").Methods(http.MethodPost).Handler(&loginHandler{DB: db})
	router.Path("/api/auth/login").Methods(http.MethodDelete).Handler(&logoutHandler{DB: db})
	router.Path("/api/auth/user").Methods(http.MethodGet).Handler(&fetchUsersHandler{DB: db})

	router.Path("/api/user").Methods(http.MethodPost).Handler(&registerUserHandler{DB: db})
	// router.
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOK, originsOK, methodsOK, credsOK)(router)))
}
