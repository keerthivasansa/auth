package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func validate(provider Provider, providerId string, password string) (User, bool) {
	pwdHash := getHash(password)
	user, err := getUserWithProvider(provider, providerId)
	if err != nil {
		panic(err)
	}
	if user.Password == pwdHash {
		return user, true
	} else {
		return User{}, false
	}
}

func loginToJwt(provider Provider, providerId string, password string) (string, error) {
	user, valid := validate(provider, providerId, password)
	if !valid {
		return "", errors.New("credentials doesn't match")
	}
	token := generateJWT(JWTClaims{Data: UserData{UserId: user.Id}, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour)),
	}})
	return token, nil
}
