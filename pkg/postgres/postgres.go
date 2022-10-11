package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"login-task/pkg/configs"
)

func OpenDBConn() *sql.DB {
	config := configs.GetConfig()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
