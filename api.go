package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	JWT string `json:"jwt"`
}

func signInRoute(c *gin.Context) {
	var body User
	c.Bind(&body)
	jwt, err := loginToJwt(body.Provider, body.ProviderId, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to login user",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		JWT: jwt,
	})
}

func signUpRoute(c *gin.Context) {
	var body User
	c.Bind(&body)
	usr, err := createUser(body.Provider, body.ProviderId, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, usr)
}

func verifyRoute(c *gin.Context) {
	prefix := "Bearer "
	authHeader := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, prefix)
	_, claims, err := verifyJWT(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to validate and get JWT claims."})
		return
	}
	c.JSON(http.StatusOK, claims)
}
