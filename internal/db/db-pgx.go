//go:build !debug
// +build !debug

package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Init(dsn string) {
	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("cannot open database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}
}
