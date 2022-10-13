package storage

import (
	"log"
	"login-task/pkg/postgres"
)

func GetUser(login, password string) (string, string, error) {
	db := postgres.OpenDBConn()
	var mail, pass string
	if err := db.Get(&mail, "SELECT login FROM users WHERE login = $1", login); err != nil {
		return "", "", err
	}
	if err := db.Get(&pass, "SELECT password FROM users WHERE password = $1", password); err != nil {
		return "", "", err
	}
	return mail, pass, nil
}

func CreateUser(login, password string) {
	db := postgres.OpenDBConn()
	_, err := db.Exec("INSERT INTO users (login, password) VALUES ($1,$2)", login, password)
	if err != nil {
		log.Fatal("Insert data error!")
	}
}
