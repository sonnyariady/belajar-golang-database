package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenKoneksi(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajargolangudemy")
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi OK")
	defer db.Close()
}
