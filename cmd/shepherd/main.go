package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
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

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("USERS"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("SESSION"))
		if err != nil {
			return fmt.Errorf("could not create days bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func registerUser(db *bolt.DB, user *User) error {
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(user); err != nil {
		return fmt.Errorf("unable to encode the user data: %w", err)
	}
	data := buff.Bytes()
	log.Printf("Data: %s", string(data))
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("USERS")).Put([]byte(user.Email), data)
		if err != nil {
			return fmt.Errorf("could not put user data: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to update the data when registering the user: %w", err)
	}
	log.Printf("User data was registerred")
	return nil
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalf("unable to create db")
	}
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

	router.Path("/api/user").Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr := &User{}
		if err := json.NewDecoder(r.Body).Decode(usr); err != nil {
			http.Error(w, "Unable to parse the request", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		registerUser(db, usr)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Successfully Registered")
	})
	// router.
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOK, originsOK, methodsOK, credsOK)(router)))
}
