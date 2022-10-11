package storage

import (
	"log"
	"login-task/pkg/postgres"
)

func GetUser(login, password string) {

}

func CreateUser(login, password string) {
	db := postgres.OpenDBConn()
	_, err := db.Exec("INSERT INTO users (login, password) VALUES ($1,$2)", login, password)
	if err != nil {
		log.Fatal("Insert data error!")
	}
}
