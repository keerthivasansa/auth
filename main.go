package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	r.POST("/auth/signin", signInRoute)
	r.POST("/auth/signup", signUpRoute)
	r.GET("/auth/verify", verifyRoute)

	log.Fatal(r.Run(":5000"))
}
