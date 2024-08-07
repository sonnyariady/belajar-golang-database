package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestKaryawanInsert(t *testing.T) {
	karyawanRepository := NewKaryawanRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	karyawan := entity.KaryawanTable{
		nama:         "Jaka",
		jabatan:      "Accounting",
		married:      true,
		tanggallahir: sql.NullTime{Time: time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{String: "Henny", Valid: true},
		gaji:         sql.NullInt64{Int64: 15, Valid: true},
	}
	hasil, err := karyawanRepository.Insert(ctx, karyawan)
	if err != nil {
		panic(err)
	}
	fmt.Println(hasil)
}
