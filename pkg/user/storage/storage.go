package storage

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"login-task/pkg/postgres"
)

func GetUser(login, password string) (string, error) {
	var db *sqlx.DB
	var err error
	if db, err = postgres.OpenDBConn(); err != nil {
		return "", errors.New("DB connection error")
	}
	defer db.Close()
	var id string
	if err := db.Get(&id, "SELECT id FROM users WHERE login = $1 AND password=$2", login, password); err != nil {
		return "", err
	}
	return id, nil
}

func CreateUser(login, password string) error {
	var db *sqlx.DB
	var err error
	if db, err = postgres.OpenDBConn(); err != nil {
		return errors.New("DB connection error")
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (login, password) VALUES ($1,$2)", login, password)
	if err != nil {
		return err
	}
	return nil
}

func AddRefreshToken(token, id string) error {
	var db *sqlx.DB
	var err error
	if db, err = postgres.OpenDBConn(); err != nil {
		return errors.New("DB connection error")
	}
	defer db.Close()
	_, err = db.Exec("UPDATE users SET refreshtoken = $1 WHERE id = $2", token, id)
	if err != nil {
		return err
	}
	return nil
}

func GetRefreshToken(id string) (string, error) {
	var db *sqlx.DB
	var err error
	if db, err = postgres.OpenDBConn(); err != nil {
		return "", errors.New("DB connection error")
	}
	defer db.Close()
	var tokenVar string
	if err = db.Get(&tokenVar, "SELECT refreshtoken FROM users WHERE id=$1", id); err != nil {
		return "", err
	}
	return tokenVar, nil
}
