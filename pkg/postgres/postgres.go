package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"login-task/pkg/configs"
)

func OpenDBConn() *sqlx.DB {
	config := configs.GetConfig()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
