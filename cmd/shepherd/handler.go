package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
)

type loginHandler struct {
	DB *bolt.DB
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	usr := &User{}
	err := json.NewDecoder(r.Body).Decode(usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	expectedPasswd, err := fetchPassword(h.DB, usr.Email)
	if err != nil {
		http.Error(w, "unable to fetch password for the user", http.StatusInternalServerError)
		return
	}
	if usr.Password != expectedPasswd {
		http.Error(w, "unauthorized user", http.StatusUnauthorized)
		return
	}
	u, err := fetchUser(h.DB, usr.Email)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	u.Password = "*******"
	token, err := createToken(h.DB, u)
	if err != nil {
		http.Error(w, "unable to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := &tokenResponseObj{
		Token: token,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Unknown error", http.StatusInternalServerError)
		return
	}
	return
}

type logoutHandler struct {
	DB *bolt.DB
}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
}

type fetchUsersHandler struct {
	DB *bolt.DB
}

func (h *fetchUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
}

type registerUserHandler struct {
	DB *bolt.DB
}

func (h *registerUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	usr := &User{}
	if err := json.NewDecoder(r.Body).Decode(usr); err != nil {
		http.Error(w, "Unable to parse the request", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	registerUser(h.DB, usr)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully Registered")
}
