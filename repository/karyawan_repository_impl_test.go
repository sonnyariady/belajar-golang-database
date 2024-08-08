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
		Nama:         "Jaka",
		Jabatan:      "Accounting",
		Married:      true,
		Tanggallahir: sql.NullTime{Time: time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local), Valid: true},
		Pasangan:     sql.NullString{String: "Henny", Valid: true},
		Gaji:         sql.NullInt64{Int64: 15, Valid: true},
	}
	hasil, err := karyawanRepository.Insert(ctx, karyawan)
	if err != nil {
		panic(err)
	}
	fmt.Println(hasil)
}

func TestFindById(t *testing.T) {
	karyawanRepository := NewKaryawanRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	hasil, err := karyawanRepository.FindById(ctx, 10)
	if err != nil {
		panic(err)
	}

	fmt.Println(hasil)
}

func TestFindAll(t *testing.T) {
	karyawanRepository := NewKaryawanRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	hasil, err := karyawanRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, kar := range hasil {
		fmt.Println(kar)
	}
}
