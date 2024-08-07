package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestEmptyMSSQL(t *testing.T) {
	fmt.Println("Test driver sql server")
}

// TODO: untuk dicoba kemudian
func TestOpenKoneksiSQLServer(t *testing.T) {

	connString := "Server=localhost;Database=Kampus;Integrated Security=SSPI;"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Verify connection
	err = db.Ping()
	if err != nil {
		panic("Error pinging database: " + err.Error())
	}

	fmt.Println("Connected to the database successfully")

	fmt.Println("Koneksi OK")

}
