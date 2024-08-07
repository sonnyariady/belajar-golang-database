package belajar_golang_database

import (
	"database/sql"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetConnectionSQLServer() *sql.DB {
	connString := "Server=localhost;Database=Kampus;Integrated Security=SSPI;"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
