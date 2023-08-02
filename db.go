package main

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

func getDatabase() *sql.DB {
	url := os.Getenv("DB_URL")
	if _db != nil {
		return _db
	}
	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	_db = conn
	return _db
}

func insertUser(user User) error {
	db := getDatabase()
	res, err := db.Exec("INSERT INTO users (id, password, provider, provider_id, created_at) VALUES (?, ?, ?, ?, ?)", user.Id, user.Password, user.Provider, user.ProviderId, time.Now())
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("failed to insert row")
	}
	return nil
}

func getUserWithProvider(provider Provider, providerId string) (User, error) {
	var user User
	db := getDatabase()
	row := db.QueryRow("SELECT id, password, provider, provider_id, created_at FROM users WHERE provider = ? AND provider_id = ?", provider, providerId)
	err := row.Scan(&user.Id, &user.Password, &user.Provider, &user.ProviderId, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
