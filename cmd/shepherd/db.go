package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
)

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

func fetchPassword(db *bolt.DB, email string) (string, error) {
	var password string
	err := db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte("DB")).Bucket([]byte("USERS")).Get([]byte(email))
		u := &User{}
		err := gob.NewDecoder(bytes.NewReader(data)).Decode(u)
		if err != nil {
			return fmt.Errorf("unable to decode the data retrieved from the db for the user with email %s: %w", email, err)
		}
		log.Printf("Retrieving Password for the user with email ID %s", email)
		password = u.Password
		return nil
	})
	if err != nil {
		return "", err
	}
	return password, nil

}

func registerUser(db *bolt.DB, user *User) error {
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(user); err != nil {
		return fmt.Errorf("unable to encode the user data: %w", err)
	}
	data := buff.Bytes()
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

func fetchUser(db *bolt.DB, email string) (*User, error) {
	usr := &User{}
	err := db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte("DB")).Bucket([]byte("USERS")).Get([]byte(email))
		err := gob.NewDecoder(bytes.NewReader(data)).Decode(usr)
		if err != nil {
			return fmt.Errorf("unable to decode the data retrieved from the db for the user with email %s: %w", email, err)
		}
		log.Printf("Retrieving Password for the user with email ID %s", email)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func createToken(db *bolt.DB, user *User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("unable to sign the token: %w", err)
	}

	return tokenString, nil
}
