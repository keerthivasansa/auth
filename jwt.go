package main

import (
	"crypto/hmac"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	UserId string `json:"userId"`
}
type JWTClaims struct {
	Data UserData `json:"data"`
	jwt.RegisteredClaims
}

func getKey() []byte {
	keyHmac := hmac.New(jwt.SigningMethodHS256.Hash.New, []byte("Hello world"))
	keyBytes := keyHmac.Sum(nil)
	return keyBytes
}

func generateJWT(claims JWTClaims) string {
	key := getKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	return tokenStr
}

func verifyJWT(tokenStr string) (*jwt.Token, JWTClaims, error) {
	key := getKey()
	var claims JWTClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid algorithm: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, JWTClaims{}, err
	}
	return token, claims, nil
}
