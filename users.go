package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func getHash(data string) string {
	hash := sha256.New().Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func createUser(provider Provider, providerId string, password string) (User, error) {
	pwdHash := getHash(password)
	id := uuid.NewString()
	user := User{
		Id:         id,
		Password:   pwdHash,
		ProviderId: providerId,
		Provider:   provider,
		CreatedAt: time.Now(),
	}
	err := insertUser(user)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}
