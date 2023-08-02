package main

import "time"

type Provider string

const (
	GoogleProvider Provider = "google"
	EmailProvider  Provider = "email"
	GithubProvider Provider = "github"
)

type User struct {
	Id         string
	Password   string
	ProviderId string `json:"providerId"`
	Provider   Provider
	CreatedAt  time.Time
}
