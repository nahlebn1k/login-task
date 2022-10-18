package storage

import (
	"log"
	"login-task/pkg/postgres"
)

func GetUser(login, password string) (string, error) {
	db := postgres.OpenDBConn()
	var id string
	if err := db.Get(&id, "SELECT id FROM users WHERE login = $1 AND password=$2", login, password); err != nil {
		return "", err
	}
	return id, nil
}

func CreateUser(login, password string) {
	db := postgres.OpenDBConn()
	_, err := db.Exec("INSERT INTO users (login, password) VALUES ($1,$2)", login, password)
	if err != nil {
		log.Fatal("Insert data error!")
	}
}

func AddRefreshToken(token, id string) {
	db := postgres.OpenDBConn()
	_, err := db.Exec("UPDATE users SET refreshtoken = $1 WHERE id = $2", token, id)
	if err != nil {
		log.Fatal("insert data error!")
	}
}

func GetRefreshToken(id string) (string, error) {
	db := postgres.OpenDBConn()
	var tokenVar string
	if err := db.Get(&tokenVar, "SELECT refreshtoken FROM users WHERE id=$1", id); err != nil {
		return "", err
	}
	return tokenVar, nil
}
