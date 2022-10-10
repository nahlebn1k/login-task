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
		config.DB.User, config.DB.Pass, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
