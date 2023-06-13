package helper

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {

	// SET DATABASE MENGGUNAKAN PASSWORD
	// sql.Open("mysql", "mysql:password@tcp(localhost:3306)/test")

	// SET DATABASE TANPA DATABASE
	db, err := sql.Open("mysql", "mysql@tcp(localhost:3306)/test")
	PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
