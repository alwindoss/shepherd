package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func extractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		log.Printf("no token found")
		return "", fmt.Errorf("no token found")
	}
	bearerTokenSlice := strings.Split(bearerToken, " ")
	log.Printf("DEBUG::Length of Token Header:  %d", len(bearerTokenSlice))
	if len(bearerTokenSlice) != 2 {
		log.Printf("token malformed")
		return "", fmt.Errorf("token malformed")
	}
	log.Printf("DEBUG::Bearer Value: %s", bearerTokenSlice[0])
	if strings.ToUpper(bearerTokenSlice[0]) != "BEARER" {
		log.Printf("token malformed")
		return "", fmt.Errorf("token malformed")
	}
	token := bearerTokenSlice[1]
	return token, nil
}

func authenticate(r *http.Request) (*User, *Claims, error) {
	token, err := extractToken(r)
	if err != nil {
		log.Printf("unable to extract token")
		return nil, nil, err
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Printf("unable to parse with claims")
		if err == jwt.ErrSignatureInvalid {
			return nil, nil, err
		}
		return nil, nil, err
	}
	if !tkn.Valid {
		log.Printf("invalid token")
		return nil, nil, err
	}
	return claims.User, claims, nil
}

func refresToken(r *http.Request) {

}
